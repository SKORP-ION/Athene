package service

import (
	"Athena/history"
	"Athena/history/api"
	. "Athena/log"
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strings"
)

type GrpcService struct {
	api.UnimplementedActionsAndHistoryServer
}

func (GrpcService) GetHistory(c context.Context, hr *api.HistoryRequest) (*api.HistoryResponse, error) {
	records, err := history.GetHistory(hr.Id)

	if err != nil {
		Error.Printf("Failed to get history for property %d. Error: %d\n", hr.Id, err)
		return nil, err
	}

	result := api.HistoryResponse{}

	for _, rec := range records {
		result.Records = append(result.Records, &api.Record{
			Action: rec.Action,
			Note:   rec.Note,
			Date:   timestamppb.New(rec.Date),
			User:   rec.StringUser,
		})
	}

	Info.Printf("Send history for property %d\n", hr.Id)
	return &result, nil
}

func (GrpcService) CreateCard(c context.Context, ar *api.ActionRequest) (*api.StatusWithProperties, error) {
	user := ar.User
	note := ar.Note
	properties := ar.Properties

	modelProperties := make([]history.Property, 0)

	for _, prop := range properties {
		modelProperties = append(modelProperties, history.Property{
			Inventory: strings.ToUpper(prop.Inventory),
			Serial:    strings.ToUpper(prop.Serial),
			Name:      prop.Name,
			State:     1,
		})
	}

	existsProperties, err := history.CreateCard(user, note, modelProperties)

	if err != nil {
		Error.Printf("Failed to create card. Error: %s\n", err)
		return &api.StatusWithProperties{
			Ok:      false,
			Message: "Internal error",
		}, err
	}

	resultProp := make([]*api.ExistsProperty, 0)

	for _, prop := range existsProperties {
		resultProp = append(resultProp, &api.ExistsProperty{
			Id:        prop.Id,
			Serial:    prop.Serial,
			Inventory: prop.Inventory,
			Name:      prop.Name,
		})
	}

	Info.Printf("Successfully created cards\n")
	return &api.StatusWithProperties{
		Ok:         true,
		Message:    "Success",
		Properties: resultProp,
	}, nil
}

func (GrpcService) InstallOnWorkspace(c context.Context, ar *api.ActionRequest) (*api.Status, error) {
	propertiesModel := make([]*history.Property, 0)

	for _, prop := range ar.Properties {
		propertiesModel = append(propertiesModel, &history.Property{
			Id:        prop.Id,
			Serial:    prop.Serial,
			Inventory: prop.Inventory,
			Name:      prop.Name,
			State:     5,
		})
	}

	err := history.DoAction(ar.User, ar.Note, propertiesModel, 5, true)

	if err != nil {
		Error.Printf("Failed to install on workspace. Error: %s\n", err)
		return &api.Status{
			Ok:      false,
			Message: "Internal error",
		}, err
	}

	Info.Printf("Successfully installed on workspace\n")
	return &api.Status{
		Ok:      true,
		Message: "Success",
	}, nil
}

func (GrpcService) RemoveFromWorkspace(c context.Context, ar *api.ActionRequest) (*api.Status, error) {
	propertiesModel := make([]*history.Property, 0)

	for _, prop := range ar.Properties {
		propertiesModel = append(propertiesModel, &history.Property{
			Id:        prop.Id,
			Serial:    prop.Serial,
			Inventory: prop.Inventory,
			Name:      prop.Name,
			State:     6,
		})
	}

	err := history.DoAction(ar.User, ar.Note, propertiesModel, 6, true)

	if err != nil {
		Error.Printf("Failed to remove from workspace. Error: %s\n", err)
		return &api.Status{
			Ok:      false,
			Message: "Internal error",
		}, err
	}

	Info.Printf("Successfully removed from workspace\n")
	return &api.Status{
		Ok:      true,
		Message: "Success",
	}, nil
}

func (GrpcService) NeedRepair(c context.Context, ar *api.ActionRequest) (*api.Status, error) {
	propertiesModel := make([]*history.Property, 0)

	for _, prop := range ar.Properties {
		propertiesModel = append(propertiesModel, &history.Property{
			Id:        prop.Id,
			Serial:    prop.Serial,
			Inventory: prop.Inventory,
			Name:      prop.Name,
			State:     7,
		})
	}

	err := history.DoAction(ar.User, ar.Note, propertiesModel, 7, true)

	if err != nil {
		Error.Printf("Failed to set status \"Need repair\". Error: %s\n", err)
		return &api.Status{
			Ok:      false,
			Message: "Internal error",
		}, err
	}

	Info.Printf("Successfully set status \"Need repair\"\n")
	return &api.Status{
		Ok:      true,
		Message: "Success",
	}, nil
}

func (GrpcService) SendToRepair(c context.Context, ar *api.ActionRequest) (*api.Status, error) {
	propertiesModel := make([]*history.Property, 0)

	for _, prop := range ar.Properties {
		propertiesModel = append(propertiesModel, &history.Property{
			Id:        prop.Id,
			Serial:    prop.Serial,
			Inventory: prop.Inventory,
			Name:      prop.Name,
			State:     8,
		})
	}

	err := history.DoAction(ar.User, ar.Note, propertiesModel, 8, true)

	if err != nil {
		Error.Printf("Failed to send to repair. Error: %d\n", err)
		return &api.Status{
			Ok:      false,
			Message: "Internal error",
		}, err
	}

	Info.Printf("Successfully sent to repair\n")
	return &api.Status{
		Ok:      true,
		Message: "Success",
	}, nil
}

func (GrpcService) ReceiveFromRepair(c context.Context, ar *api.ActionRequest) (*api.Status, error) {
	propertiesModel := make([]*history.Property, 0)

	for _, prop := range ar.Properties {
		propertiesModel = append(propertiesModel, &history.Property{
			Id:        prop.Id,
			Serial:    prop.Serial,
			Inventory: prop.Inventory,
			Name:      prop.Name,
			State:     9,
		})
	}

	err := history.DoAction(ar.User, ar.Note, propertiesModel, 9, true)

	if err != nil {
		Error.Printf("Failed to receive from repair. Error: %s\n", err)
		return &api.Status{
			Ok:      false,
			Message: "Internal error",
		}, err
	}

	Info.Printf("Successfully receive from repair\n")
	return &api.Status{
		Ok:      true,
		Message: "Success",
	}, nil
}

func (GrpcService) Archive(c context.Context, ar *api.ActionRequest) (*api.Status, error) {
	propertiesModel := make([]*history.Property, 0)

	for _, prop := range ar.Properties {
		propertiesModel = append(propertiesModel, &history.Property{
			Id:        prop.Id,
			Serial:    prop.Serial,
			Inventory: prop.Inventory,
			Name:      prop.Name,
			State:     11,
		})
	}

	err := history.DoAction(ar.User, ar.Note, propertiesModel, 11, true)

	if err != nil {
		Error.Printf("Failed to archive property. Error: %s\n", err)
		return &api.Status{
			Ok:      false,
			Message: "Internal error",
		}, err
	}

	Info.Println("Successfully archive property")
	return &api.Status{
		Ok:      true,
		Message: "Success",
	}, nil
}

func (GrpcService) DeArchive(c context.Context, ar *api.ActionRequest) (*api.Status, error) {
	propertiesModel := make([]*history.Property, 0)

	for _, prop := range ar.Properties {
		propertiesModel = append(propertiesModel, &history.Property{
			Id:        prop.Id,
			Serial:    prop.Serial,
			Inventory: prop.Inventory,
			Name:      prop.Name,
			State:     14,
		})
	}

	err := history.DoAction(ar.User, ar.Note, propertiesModel, 14, true)

	if err != nil {
		Error.Printf("Failed to remove from archive properties. Error: %s\n", err)
		return &api.Status{
			Ok:      false,
			Message: "Internal error",
		}, err
	}

	Info.Println("Successfully removed from archive")
	return &api.Status{
		Ok:      true,
		Message: "Success",
	}, nil
}

func (GrpcService) ChangeName(c context.Context, ar *api.ActionRequest) (*api.Status, error) {
	propertiesModel := make([]*history.Property, 0)

	for _, prop := range ar.Properties {
		propertiesModel = append(propertiesModel, &history.Property{
			Id:        prop.Id,
			Serial:    prop.Serial,
			Inventory: prop.Inventory,
			Name:      prop.Name,
			State:     13,
		})
	}

	for _, prop := range propertiesModel {
		err := history.ChangeField(prop, "name")

		if err != nil {
			Error.Printf("Failed to change name %s for prop %d. Error: %s\n", prop.Name, prop.Id, err)
			return &api.Status{
				Ok:      false,
				Message: "Internal error",
			}, err
		}
	}

	err := history.DoAction(ar.User, ar.Note, propertiesModel, 13, false)

	if err != nil {
		Error.Printf("Failed to change name record. Error: %s", err)
		return &api.Status{
			Ok:      false,
			Message: "Internal error",
		}, err
	}

	Info.Println("Successfully change name for properties")
	return &api.Status{
		Ok:      true,
		Message: "Success",
	}, nil
}

func (GrpcService) ChangeInventory(c context.Context, ar *api.ActionRequest) (*api.Status, error) {
	propertiesModel := make([]*history.Property, 0)

	for _, prop := range ar.Properties {
		propertiesModel = append(propertiesModel, &history.Property{
			Id:        prop.Id,
			Serial:    prop.Serial,
			Inventory: prop.Inventory,
			Name:      prop.Name,
			State:     12,
		})
	}

	for _, prop := range propertiesModel {
		err := history.ChangeField(prop, "inventory")

		if err != nil {
			Error.Printf("Failed to change inventory %s for prop %d. Error: %s\n", prop.Inventory, prop.Id, err)
			return &api.Status{
				Ok:      false,
				Message: "Internal error",
			}, err
		}
	}

	err := history.DoAction(ar.User, ar.Note, propertiesModel, 12, false)

	if err != nil {
		Error.Printf("Failed to create change inventory record. Error: %s", err)
		return &api.Status{
			Ok:      false,
			Message: "Internal error",
		}, err
	}

	Info.Println("Successfully change inventory for properties")
	return &api.Status{
		Ok:      true,
		Message: "Success",
	}, nil
}

func (GrpcService) AddToGroup(c context.Context, ar *api.ActionRequest) (*api.Status, error) {
	propertiesModel := make([]*history.Property, 0)

	group, err := history.GetGroup(ar.SearchId)

	if err != nil {
		Error.Printf("Failed to get group by id %d. Error: %s\n", ar.SearchId, err)
		return &api.Status{
			Ok:      false,
			Message: "Internal error",
		}, err
	}

	for _, prop := range ar.Properties {
		propertiesModel = append(propertiesModel, &history.Property{
			Id:    prop.Id,
			State: 15,
		})
	}

	note := fmt.Sprintf("Добавлено в группу \"%s\". Заметка: %s",
		group.Name, ar.Note)

	err = history.DoAction(ar.User, note, propertiesModel, 15, false)

	if err != nil {
		Error.Printf("Failed to add properties to group %s. Error: %s\n", group.Name, err)
		return &api.Status{
			Ok:      false,
			Message: "Internal error",
		}, err
	}

	Info.Printf("Successfully added properties to group %s\n", group.Name)
	return &api.Status{
		Ok:      true,
		Message: "Success",
	}, nil
}

func (GrpcService) RemoveFromGroup(c context.Context, ar *api.ActionRequest) (*api.Status, error) {
	propertiesModel := make([]*history.Property, 0)

	group, err := history.GetGroup(ar.SearchId)

	if err != nil {
		Error.Printf("Failed to get group by id %d. Error: %s\n", ar.SearchId, err)
		return &api.Status{
			Ok:      false,
			Message: "Internal error",
		}, err
	}

	for _, prop := range ar.Properties {
		propertiesModel = append(propertiesModel, &history.Property{
			Id:    prop.Id,
			State: 16,
		})
	}

	note := fmt.Sprintf("Удалено из группы \"%s\". Заметка: %s",
		group.Name, ar.Note)

	err = history.DoAction(ar.User, note, propertiesModel, 16, false)

	if err != nil {
		Error.Printf("Failed to remove from group %s properties. Error: %s\n", group.Name, err)
		return &api.Status{
			Ok:      false,
			Message: "Internal error",
		}, err
	}

	Info.Printf("Successfully removed from group %s\n", group.Name)
	return &api.Status{
		Ok:      true,
		Message: "Success",
	}, nil
}

func (GrpcService) SendToWarehouse(c context.Context, ar *api.ActionRequest) (*api.Status, error) {
	propertiesModel := make([]*history.Property, 0)

	warehouse, err := history.GetWarehouse(ar.SearchId)

	if err != nil {
		Error.Printf("Failed to send get warehouse by id %d. Error: %s\n", ar.SearchId, err)
		return &api.Status{
			Ok:      false,
			Message: "Internal error",
		}, err
	}

	for _, prop := range ar.Properties {
		propertiesModel = append(propertiesModel, &history.Property{
			Id:    prop.Id,
			State: 10,
		})
	}

	note := fmt.Sprintf("Перемещено на склад \"%s\". Заметка: %s",
		warehouse.Name, ar.Note)

	err = history.DoAction(ar.User, note, propertiesModel, 10, true)

	if err != nil {
		Error.Printf("Failed to create send to warehouse %s record. Error: %s\n", warehouse.Name, err)
		return &api.Status{
			Ok:      false,
			Message: "Internal error",
		}, err
	}

	Info.Printf("Successfully sent to warehouse %s\n", warehouse.Name)
	return &api.Status{
		Ok:      true,
		Message: "Success",
	}, nil
}

func (GrpcService) GiveToEmployee(c context.Context, ar *api.ActionRequest) (*api.Status, error) {
	propertiesModel := make([]*history.Property, 0)

	employee, err := history.GetEmployee(ar.SearchId)

	if err != nil {
		Error.Printf("Failed to get employee by id %d. Error: %s\n", ar.SearchId, err)
		return &api.Status{
			Ok:      false,
			Message: "Internal error",
		}, err
	}

	for _, prop := range ar.Properties {
		propertiesModel = append(propertiesModel, &history.Property{
			Id:    prop.Id,
			State: 3,
		})
	}

	note := fmt.Sprintf("Выдано пользователю \"%s\". Заметка: %s",
		employee.Name, ar.Note)

	err = history.DoAction(ar.User, note, propertiesModel, 3, true)

	if err != nil {
		Error.Printf("Failed to create give to employee %s record. Error: %s\n", employee.Name, err)
		return &api.Status{
			Ok:      false,
			Message: "Internal error",
		}, err
	}

	Info.Printf("Successfully given to employee %s properties\n", employee.Name)
	return &api.Status{
		Ok:      true,
		Message: "Success",
	}, nil
}

func (GrpcService) TakeFromEmployee(c context.Context, ar *api.ActionRequest) (*api.Status, error) {
	propertiesModel := make([]*history.Property, 0)

	employee, err := history.GetEmployee(ar.SearchId)

	if err != nil {
		Error.Printf("Failed to get employee by id %d. Error: %s\n", ar.SearchId, err)
		return &api.Status{
			Ok:      false,
			Message: "Internal error",
		}, err
	}

	for _, prop := range ar.Properties {
		propertiesModel = append(propertiesModel, &history.Property{
			Id:    prop.Id,
			State: 4,
		})
	}

	note := fmt.Sprintf("Получено от пользователя \"%s\". Заметка: %s",
		employee.Name, ar.Note)

	err = history.DoAction(ar.User, note, propertiesModel, 4, true)

	if err != nil {
		Info.Printf("Failed to create take from employee %s record. Error: %s\n", employee.Name, err)
		return &api.Status{
			Ok:      false,
			Message: "Internal error",
		}, err
	}

	Info.Printf("Successfully taken from employee %s\n", employee.Name)
	return &api.Status{
		Ok:      true,
		Message: "Success",
	}, nil
}
