package service

import (
	"Athena/groups"
	"Athena/groups/api"
	. "Athena/log"
	"context"
	"errors"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

type GrpcService struct {
	api.UnimplementedGroupsServer
}

func (GrpcService) CreateGroup(c context.Context, gr *api.Group) (*api.GroupStatus, error) {
	model := groups.Group{
		Name:        gr.Name,
		Description: gr.Description,
	}
	err := groups.CreateGroup(&model, gr.WhoDisplayName)

	if err != nil {
		Error.Println("Failed to create group. Error: ", err)
		return &api.GroupStatus{
			Ok:      false,
			Message: "Internal error",
		}, err
	}

	Info.Printf("Successfully created group %s\f", model.Name)
	return &api.GroupStatus{
		Ok:      true,
		Message: "Success",
	}, nil
}

func (GrpcService) RemoveGroup(c context.Context, gr *api.Group) (*api.GroupStatus, error) {
	model := groups.Group{
		Id: gr.Id,
	}

	err := groups.RemoveGroup(&model)

	if err != nil {
		Error.Printf("Failed to remove group. id:%d. Error: %s\n", model.Id, err)
		return &api.GroupStatus{
			Ok:      false,
			Message: "Internal error",
		}, err
	}

	Info.Printf("Successfully removed group with id: %d\n", model.Id)
	return &api.GroupStatus{
		Ok:      true,
		Message: "Success",
	}, nil
}

func (GrpcService) GetGroups(c context.Context, gr *api.GetGroupsReq) (*api.GetGroupsRes, error) {
	name := gr.Name

	grps, err := groups.GetGroups(name)

	if err != nil {
		Error.Printf("Failed to get groups list by Name %s. Error: %s\n", name, err)
		return nil, err
	}

	result := &api.GetGroupsRes{}

	for _, grp := range grps {
		result.Groups = append(result.Groups, &api.Group{
			Id:             grp.Id,
			Name:           grp.Name,
			Description:    grp.Description,
			CreatedAt:      timestamppb.New(grp.CreatedAt),
			WhoDisplayName: grp.Display,
		})
	}

	Info.Printf("Successfully send groups list by name %s\n", name)
	return result, nil
}

func (GrpcService) AddPropToGroup(c context.Context, pr *api.PropReq) (*api.GroupStatus, error) {
	props := make([]*groups.PropGroup, 0)

	for _, id := range pr.PropIds {
		props = append(props, &groups.PropGroup{
			GroupId:    pr.GroupId,
			PropertyId: id,
		})
	}

	err := groups.AddToGroup(props)

	if err != nil {
		Error.Printf("Failed to add props to group. GroupId: %d. Error: %s", pr.GroupId, err)
		return &api.GroupStatus{
			Ok:      false,
			Message: "Internal error",
		}, err
	}

	Info.Printf("Successfully added properties to group %d\n", pr.GroupId)
	return &api.GroupStatus{
		Ok:      true,
		Message: "Success",
	}, err
}

func (GrpcService) RemovePropFromGroup(c context.Context, pr *api.PropReq) (*api.GroupStatus, error) {
	props := make([]*groups.PropGroup, 0)

	for _, id := range pr.PropIds {
		props = append(props, &groups.PropGroup{
			GroupId:    pr.GroupId,
			PropertyId: id,
		})
	}

	err := groups.RemoveFromGroup(props)

	if err != nil {
		Error.Printf("Failed to remove properties from group %d. Error: %s\n", pr.GroupId, err)
		return &api.GroupStatus{
			Ok:      false,
			Message: "Internal error",
		}, err
	}

	Info.Printf("Successfully removed properties from group %d\n", pr.GroupId)
	return &api.GroupStatus{
		Ok:      true,
		Message: "Success",
	}, err
}

func (GrpcService) IsInGroup(c context.Context, pr *api.PropReq) (*api.GroupStatus, error) {
	prop := &groups.PropGroup{
		GroupId:    pr.GroupId,
		PropertyId: pr.PropIds[0],
	}

	prop, err := groups.IsInGroup(prop)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		Info.Printf("Property %d not found in group %d\n", prop.PropertyId, prop.GroupId)
		return &api.GroupStatus{
			Ok:      false,
			Message: "Record not found",
		}, nil
	} else if err != nil {
		Error.Printf("Failed to check property %d in group %d. Error: %s\n", prop.PropertyId, prop.GroupId, err)
		return &api.GroupStatus{
			Ok:      false,
			Message: "Internal error",
		}, err
	}

	Info.Printf("Property %d is in group %d\n", prop.PropertyId, prop.GroupId)
	return &api.GroupStatus{
		Ok:      true,
		Message: "Record found",
	}, nil
}
