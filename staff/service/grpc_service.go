package service

import (
	. "Athena/log"
	"Athena/staff"
	"Athena/staff/api"
	"context"
	"errors"
	"fmt"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

type GrpcService struct {
	api.UnimplementedStaffServer
}

func (GrpcService) GetStaff(c context.Context, req *api.GetReq) (*api.GetResp, error) {
	params := staff.SearchStaff{
		Search:      req.Search,
		TableNumber: req.TableNumber,
		Department:  req.Department,
		Manager:     req.Manager,
		Job:         req.Job,
		Offset:      req.Offset,
		Limit:       req.Limit,
		Order:       req.Order,
	}

	employees, err := staff.GetStaff(&params)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		Info.Println("Employees not found")
		return &api.GetResp{
			Ok:    false,
			Staff: nil,
		}, nil
	} else if err != nil {
		Error.Println("Failed to get staff. Error: ", err)
		return &api.GetResp{
			Ok:    false,
			Staff: nil,
		}, err
	}

	resp := &api.GetResp{Ok: true}

	for _, emp := range employees {
		resp.Staff = append(resp.Staff, &api.StaffEmployee{
			Id:         emp.Id,
			Table:      emp.Table,
			Name:       emp.Name,
			Manager:    emp.Manager,
			Department: emp.Department,
			Job:        emp.Job,
			CreatedAt:  timestamppb.New(emp.CreatedAt),
		})
	}
	Info.Println("Send staff list")
	return resp, nil
}

func (GrpcService) GetCount(c context.Context, req *api.GetReq) (*api.CountResp, error) {
	params := staff.SearchStaff{
		Search:      req.Search,
		TableNumber: req.TableNumber,
		Department:  req.Department,
		Manager:     req.Manager,
		Job:         req.Job,
		Offset:      req.Offset,
		Limit:       req.Limit,
		Order:       req.Order,
	}

	count, err := staff.GetCount(&params)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		Info.Println("Record not found. Count = 0")
		return &api.CountResp{
			Ok:    false,
			Count: 0,
		}, nil
	} else if err != nil {
		Error.Println("Failed to get count. Error: ", err)
		return &api.CountResp{
			Ok:    false,
			Count: 0,
		}, err
	}

	Info.Println("Send count")
	return &api.CountResp{Ok: true, Count: count}, nil
}

func (GrpcService) CreateEmployee(c context.Context, req *api.StaffEmployee) (*api.StatusWithEmployee, error) {
	emp := staff.Employee{
		Table:      req.Table,
		Name:       req.Name,
		Manager:    req.Manager,
		Department: req.Department,
		Job:        req.Job,
		CreatedAt:  req.CreatedAt.AsTime(),
	}

	emp, err := staff.CreateEmployee(emp)

	if err != nil {
		Error.Printf("Failed to create employee %s\n", req.Name)
		return &api.StatusWithEmployee{
			Ok:       false,
			Employee: nil,
		}, err
	}

	status := &api.StatusWithEmployee{
		Ok: true,
		Employee: &api.StaffEmployee{
			Id:         emp.Id,
			Table:      emp.Table,
			Name:       emp.Name,
			Manager:    emp.Manager,
			Department: emp.Department,
			Job:        emp.Job,
			CreatedAt:  timestamppb.New(emp.CreatedAt),
		},
	}

	Info.Printf("Successfully created employee %s\n", req.Name)
	return status, nil
}

func (GrpcService) GetStaffProp(c context.Context, req *api.StaffEmployee) (*api.GetStaffPropResp, error) {
	emp := staff.Employee{
		Id: req.Id,
	}
	props, err := staff.GetEmployeesProp(emp)

	if err != nil {
		Error.Printf("Failed to get employee %d properties\n", emp.Id)
		return &api.GetStaffPropResp{
			Ok:    false,
			Props: nil,
		}, err
	}

	resp := &api.GetStaffPropResp{Ok: true}

	for _, prop := range props {
		resp.Props = append(resp.Props, &api.StaffProp{
			Id:        prop.Id,
			Inventory: prop.Inventory,
			Serial:    prop.Serial,
			Name:      prop.Name,
			CreatedAt: timestamppb.New(prop.CreatedAt),
			Warehouse: prop.Warehouse,
			State:     prop.State,
			Username:  prop.Username,
			GivenAt:   timestamppb.New(prop.GivenAt),
			RecordId:  prop.Record,
		})
	}

	Info.Printf("Send employee %d properties\n", emp.Id)
	return resp, nil
}

func (GrpcService) GiveToEmployee(c context.Context, req *api.GiveReq) (*api.SStatus, error) {

	user, err := staff.GetUser(req.Username)

	if err != nil {
		Error.Printf("Failed to get user by username %s\n", req.Username)
		return &api.SStatus{
			Ok:      false,
			Message: "Internal error",
		}, err
	}

	props := make([]*staff.Property, 0)

	for _, id := range req.Ids {
		props = append(props, &staff.Property{
			Id: id,
		})
	}

	err = staff.GiveToEmployee(req.EmployeeId, user.Id, props)

	if err != nil {
		Info.Printf("Failed to give properties to employee %d\n", req.EmployeeId)
		return &api.SStatus{
			Ok:      false,
			Message: "Internal error",
		}, err
	}

	Info.Printf("Successfully give properties to employee %d\n", req.EmployeeId)
	return &api.SStatus{
		Ok:      true,
		Message: "Success",
	}, nil
}

func (GrpcService) TakeFromEmployee(c context.Context, req *api.GiveReq) (*api.SStatus, error) {

	props := make([]*staff.Property, 0)

	for _, id := range req.Ids {
		props = append(props, &staff.Property{
			Id: id,
		})
	}

	err := staff.TakeFromEmployee(req.EmployeeId, props)

	if err != nil {
		Error.Printf("Failed to take properties from employee %d\n", req.EmployeeId)
		return &api.SStatus{
			Ok:      false,
			Message: "Internal error",
		}, err
	}

	Info.Println("Successfully take properties from employee %d\n", req.EmployeeId)
	return &api.SStatus{
		Ok:      true,
		Message: "Success",
	}, nil
}

func (GrpcService) IsWithEmployee(c context.Context, req *api.GiveReq) (*api.SStatus, error) {
	prop := req.Ids[0]

	record := staff.PropRecord{PropertyId: prop}

	record, err := staff.GetRecord(record)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		Info.Printf("Property %d is not with employee\n", prop)
		return &api.SStatus{
			Ok:      false,
			Message: "Record not found",
		}, nil
	} else if err != nil {
		Error.Printf("Failed to check property %d. Error: %s\n", prop, err)
		return &api.SStatus{
			Ok:      false,
			Message: "Internal error",
		}, err
	}

	if record.StaffId != req.EmployeeId {
		Info.Printf("Another employee has property %d\n", prop)
		return &api.SStatus{
			Ok:      false,
			Message: fmt.Sprintf("Last owner is %d", record.StaffId),
		}, nil
	}

	Info.Printf("Property %d now with right employee\n", prop)
	return &api.SStatus{
		Ok:      true,
		Message: "Property now with this employee",
	}, nil
}

func (GrpcService) GetEmployee(c context.Context, req *api.GetEmployeeReq) (*api.StaffEmployee, error) {
	emp, err := staff.GetEmployeeById(req.Id)

	if err != nil {
		Error.Printf("Failed to get employee by id %d. Error: %s\n", req.Id, err)
		return nil, err
	}

	Info.Printf("Send employee by id %d\n", req.Id)
	return &api.StaffEmployee{
		Id:         emp.Id,
		Name:       emp.Name,
		Job:        emp.Job,
		Manager:    emp.Manager,
		Table:      emp.Table,
		Department: emp.Department,
	}, nil
}

func (GrpcService) GetRecord(c context.Context, req *api.SRecord) (*api.GetRecordResp, error) {
	record, err := staff.GetRecordById(req.Id)

	if err != nil {
		Error.Printf("Failed to get record by id %d. Error: %s\n", req.Id, err)
		return nil, err
	}

	emp, err := staff.GetEmployeeById(record.StaffId)

	if err != nil {
		Error.Printf("Failed to get employee by id %d. Error: %s\n", record.StaffId, err)
		return nil, err
	}

	prop, err := staff.GetProperty(record.PropertyId)

	if err != nil {
		Error.Printf("Failed to get property by id %d. Error: %s\n", record.PropertyId)
		return nil, err
	}

	result := &api.GetRecordResp{
		Date: timestamppb.New(record.CreatedAt),
		Emp: &api.StaffEmployee{
			Id:         emp.Id,
			Name:       emp.Name,
			Table:      emp.Table,
			Job:        emp.Job,
			Manager:    emp.Manager,
			Department: emp.Department,
		},
		Prop: &api.StaffProp{
			Id:        prop.Id,
			Name:      prop.Name,
			Serial:    prop.Serial,
			Inventory: prop.Inventory,
		},
	}

	Info.Printf("Send record for emp %s and prop %s-%s-%s\n",
		result.Emp.Name, result.Prop.Name, result.Prop.Serial, result.Prop.Inventory)
	return result, nil
}
