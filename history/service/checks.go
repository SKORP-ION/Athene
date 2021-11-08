package service

import (
	"Athena/history"
	"Athena/history/api"
	. "Athena/log"
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
)

func (GrpcService) IsCreated(c context.Context, ps *api.PropertySerial) (*api.StatusWithProperty, error) {
	prop, err := history.GetPropertyBySerial(ps.Serial)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		Info.Printf("Property with serial %s is not created\n", ps.Serial)
		return &api.StatusWithProperty{
			Ok:      false,
			Message: "Record not found",
		}, nil
	}

	if err != nil {
		Info.Printf("Can't check is created property %s. Error: %s\n", ps.Serial, err)
		return &api.StatusWithProperty{
			Ok:      false,
			Message: "Internal Error",
		}, err
	}

	if prop.Id == 0 {
		Info.Printf("Property with serial %s is not created\n", ps.Serial)
		return &api.StatusWithProperty{
			Ok:      false,
			Message: "Property with this serial number not found",
		}, nil
	}

	Info.Printf("Property with serial %s found: %s-%s-%s\n",
		prop.Serial, prop.Name, prop.Serial, prop.Inventory)
	return &api.StatusWithProperty{
		Ok:      true,
		Message: "Success",
		Property: &api.ExistsProperty{
			Id:        prop.Id,
			Serial:    prop.Serial,
			Inventory: prop.Inventory,
			Name:      prop.Name,
		},
	}, nil
}

func (GrpcService) IsOnWorkspace(c context.Context, ps *api.PropertySerial) (*api.Status, error) {
	record, err := history.GetPropertyBySerial(ps.Serial)

	if err != nil {
		Error.Printf("Failed to get property by serial %s. Error: %s\n", ps.Serial, err)
		return &api.Status{
			Ok:      false,
			Message: "Internal Error",
		}, err
	}

	if record.State != 5 {
		Info.Printf("Property %s-%s-%s is not on workspace\n",
			record.Name, record.Serial, record.Inventory)
		return &api.Status{
			Ok:      false,
			Message: fmt.Sprintf("Last action is %d", record.State),
		}, nil
	}
	Info.Printf("Property %s-%s-%s is on workspace\n",
		record.Name, record.Serial, record.Inventory)
	return &api.Status{Ok: true, Message: "Success"}, nil
}

func (GrpcService) IsNeedsRepair(c context.Context, ps *api.PropertySerial) (*api.Status, error) {
	record, err := history.GetPropertyBySerial(ps.Serial)

	if err != nil {
		Error.Printf("Failed to get property by serial %s. Error: %s\n",
			ps.Serial, err)
		return &api.Status{
			Ok:      false,
			Message: "Internal Error",
		}, err
	}

	if record.State != 7 {
		Info.Printf("Property %s-%s-%s is no need repair.\n",
			record.Name, record.Serial, record.Inventory)
		return &api.Status{
			Ok:      false,
			Message: fmt.Sprintf("Last action is %d", record.State),
		}, nil
	}
	Info.Printf("Property %s-%s-%s is need repair.\n",
		record.Name, record.Serial, record.Inventory)
	return &api.Status{Ok: true, Message: "Success"}, nil
}

func (GrpcService) IsUnderRepair(c context.Context, ps *api.PropertySerial) (*api.Status, error) {
	record, err := history.GetPropertyBySerial(ps.Serial)

	if err != nil {
		Error.Printf("Can't get property by serial %s. Error: %s\n",
			ps.Serial, err)
		return &api.Status{
			Ok:      false,
			Message: "Internal Error",
		}, err
	}

	if record.State != 8 {
		Info.Printf("Property %s-%s-%s is not under repair.\n",
			record.Name, record.Serial, record.Inventory)
		return &api.Status{
			Ok:      false,
			Message: fmt.Sprintf("Last action is %d", record.State),
		}, nil
	}

	Info.Printf("Property %s-%s-%s is under repair\n",
		record.Name, record.Serial, record.Inventory)
	return &api.Status{Ok: true, Message: "Success"}, nil
}

func (GrpcService) IsInArchive(c context.Context, ps *api.PropertySerial) (*api.Status, error) {
	record, err := history.GetPropertyBySerial(ps.Serial)

	if err != nil {
		Error.Printf("Can't get property by serial %s. Error: %s\n",
			ps.Serial, err)
		return &api.Status{
			Ok:      false,
			Message: "Internal Error",
		}, err
	}

	if record.State != 11 {
		Info.Printf("Property %s-%s-%s is not in archive.\n",
			record.Name, record.Serial, record.Inventory)
		return &api.Status{
			Ok:      false,
			Message: fmt.Sprintf("Last action is %d", record.State),
		}, nil
	}

	Info.Printf("Property %s-%s-%s is in archive.\n",
		record.Name, record.Serial, record.Inventory)
	return &api.Status{Ok: true, Message: "Success"}, nil
}

func (GrpcService) IsInStock(c context.Context, ps *api.PropertySerial) (*api.Status, error) {
	record, err := history.GetPropertyBySerial(ps.Serial)

	if err != nil {
		Error.Printf("Can't get property by serial %s. Error: %s\n",
			ps.Serial, err)
		return &api.Status{
			Ok:      false,
			Message: "Internal Error",
		}, err
	}

	switch record.State {
	case 1, 4, 6, 9, 10, 12, 13, 14:
		{
			Info.Printf("Property %s-%s-%s is in stock\n",
				record.Name, record.Serial, record.Inventory)
			return &api.Status{Ok: true, Message: "Success"}, nil
		}
	default:
		{
			Info.Printf("Property %s-%s-%s is not in stock\n",
				record.Name, record.Serial, record.Inventory)
			return &api.Status{
				Ok:      false,
				Message: fmt.Sprintf("Last action is %d", record.State),
			}, nil
		}
	}
}
