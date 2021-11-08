package service

import (
	. "Athena/log"
	"Athena/property"
	"Athena/property/api"
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var NotInStock = map[uint32]bool{
	3:  false,
	5:  false,
	7:  false,
	8:  false,
	11: false,
}

type GrpcService struct {
	api.UnimplementedPropertyServer
}

func (GrpcService) GetProperty(c context.Context, sp *api.SearchParams) (*api.Properties, error) {
	properties, err := property.GetProperty(&property.SearchParams{
		Inventory: sp.Inventory,
		Serial:    sp.Serial,
		Name:      sp.Name,
		Action:    sp.Action,
		Warehouse: sp.Warehouse,
		Offset:    sp.Offset,
		Limit:     sp.Limit,
		Order:     sp.Order,
		Groups:    sp.Groups,
	})

	if err != nil {
		Error.Println("Failed to find property by condition")
		return nil, err
	}

	Info.Println("Send property list.")
	return fromModelToGrpcType(properties), nil

}

func (GrpcService) GetWarehouses(c context.Context, _ *api.EmptyPropertyRequest) (*api.Warehouses, error) {
	warehouses, err := property.GetWarehouses()
	if err != nil {
		Error.Println("Failed to get warehouses list.")
		return nil, err
	}

	data := make([]*api.Warehouse, 0)

	for _, value := range warehouses {
		data = append(data, &api.Warehouse{
			Id:   value.Id,
			Name: value.Name,
		})
	}
	Info.Println("Send warehouses list")
	return &api.Warehouses{Data: data}, nil
}

func (GrpcService) GetCount(c context.Context, sp *api.SearchParams) (*api.Count, error) {

	count, err := property.GetCount(&property.SearchParams{
		Inventory: sp.Inventory,
		Serial:    sp.Serial,
		Name:      sp.Name,
		Action:    sp.Action,
		Warehouse: sp.Warehouse,
		Offset:    sp.Offset,
		Limit:     sp.Limit,
		Order:     sp.Order,
		Groups:    sp.Groups,
	})

	if err != nil {
		Error.Println("Failed to get count")
		return nil, err
	}

	Info.Println("Send count by condition")
	return &api.Count{Number: uint64(count)}, nil
}

func (GrpcService) GetActionsList(c context.Context, _ *api.EmptyPropertyRequest) (*api.GetActionsListResponse, error) {
	getActionsListResponse := &api.GetActionsListResponse{}

	actions, err := property.GetActions()

	if err != nil {
		Error.Printf("Failed to get actions list. Error: %s\n", err)
		return nil, err
	}

	for _, action := range actions {
		getActionsListResponse.Actions = append(getActionsListResponse.Actions,
			&api.OneAction{
				Id:     action.Id,
				Name:   action.Name,
				Action: action.Action,
			})
	}
	Info.Println("Send actions list")
	return getActionsListResponse, nil
}

func (GrpcService) GetOneProperty(c context.Context, request *api.GetOnePropertyRequest) (
	*api.Property, error) {
	prop, err := property.GetOneProperty(request.Id)

	if err != nil {
		Error.Printf("Failed to get property %d. Error: %s\n", request.Id, err)
		return nil, err
	}

	Info.Printf("Successfully send property by id %d\n", request.Id)
	return &api.Property{
		Id:        prop.Id,
		Inventory: prop.Inventory,
		Serial:    prop.Serial,
		Name:      prop.Name,
		CreatedAt: timestamppb.New(prop.Created_at),
		UpdatedAt: timestamppb.New(prop.Updated_at),
		Warehouse: prop.Warehouse,
		State:     prop.Action,
	}, nil
}

func (GrpcService) IsOnWarehouse(c context.Context, req *api.IsInWarehouseReq) (*api.PropStatus, error) {
	prop := &property.Property{
		Id: req.PropertyId,
	}

	prop, err := property.IsOnWarehouse(prop)

	if err != nil {
		Error.Printf("Failed to check property %d is in stock\n", req.PropertyId)
		return &api.PropStatus{
			Ok:      false,
			Message: "internal error",
		}, err
	}

	if _, ok := NotInStock[prop.State]; ok {
		Info.Printf("Property %s-%s-%s is not in stock.\n", prop.Name, prop.Serial, prop.Inventory)
		return &api.PropStatus{
			Ok:      false,
			Message: "Not in stock",
		}, nil
	} else if prop.Warehouse != req.WarehouseId {
		Info.Printf("Property %s-%s-%s is not on warehouse %d.\n",
			prop.Name, prop.Serial, prop.Inventory, req.WarehouseId)
		return &api.PropStatus{
			Ok:      false,
			Message: "Record not found",
		}, nil
	}

	Info.Printf("Property %s-%s-%s is in stock\n",
		prop.Name, prop.Serial, prop.Inventory)
	return &api.PropStatus{
		Ok:      true,
		Message: "Record found",
	}, nil
}

func (GrpcService) SendToWarehouse(c context.Context, req *api.SendToWarhouseReq) (*api.PropStatus, error) {
	props := make([]*property.Property, 0)

	for _, id := range req.PropertiesId {
		props = append(props, &property.Property{
			Id:        id,
			Warehouse: req.WarehouseId,
		})
	}

	err := property.ChangeWarehouse(props)

	if err != nil {
		Error.Printf("Failed to change warehouse %d\n", req.WarehouseId)
		return &api.PropStatus{
			Ok:      false,
			Message: "Internal error",
		}, err
	}

	Info.Printf("Successfully change warehouse %d\n", req.WarehouseId)
	return &api.PropStatus{
		Ok:      true,
		Message: "Success",
	}, err
}

func fromModelToGrpcType(model []*property.Property) *api.Properties {
	properties := make([]*api.Property, 0)

	for _, prop := range model {
		properties = append(properties, &api.Property{
			Id:        prop.Id,
			Inventory: prop.Inventory,
			Serial:    prop.Serial,
			Name:      prop.Name,
			Warehouse: prop.Warehouse,
			CreatedAt: timestamppb.New(prop.Created_at),
			UpdatedAt: timestamppb.New(prop.Updated_at),
			State:     prop.Action,
		})
	}

	return &api.Properties{Properties: properties}
}
