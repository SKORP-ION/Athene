package handler

import (
	"Athena/groups/api"
	api2 "Athena/history/api"
	"Athena/http_server/gRPC/groups"
	"Athena/http_server/gRPC/history"
	"context"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
)

func Groups_GetGroups(c *gin.Context) {
	name := c.Query("name")

	grp, err := groups.Client.GetGroups(context.Background(), &api.GetGroupsReq{Name: name})

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, errors.New("internal error"))
		return
	}

	c.JSON(http.StatusOK, grp.Groups)
}

func Groups_GroupAction(c *gin.Context) {
	status := &api.GroupStatus{}
	username, err := c.Cookie("username")

	if err != nil {
		status.Ok = false
		status.Message = "Can't read username from cookie"
		c.JSON(http.StatusBadRequest, &status)
		return
	}

	group := api.Group{}

	rawData, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		status.Ok = false
		status.Message = "Can't read request body"
		c.JSON(http.StatusBadRequest, &status)
		return
	}

	err = json.Unmarshal(rawData, &group)

	if err != nil {
		status.Ok = false
		status.Message = "Bad json"
		c.JSON(http.StatusBadRequest, &status)
		return
	}

	group.WhoDisplayName = username

	switch c.FullPath() {
	case "/private/groups/CreateGroup":
		status, err = groups.Client.CreateGroup(context.Background(), &group)
	case "/private/groups/RemoveGroup":
		status, err = groups.Client.RemoveGroup(context.Background(), &group)
	}

	if err != nil {
		status = &api.GroupStatus{}
		status.Ok = false
		status.Message = "Internal error"
		c.JSON(http.StatusInternalServerError, &status)
		return
	}

	c.JSON(http.StatusOK, status)
}

func Groups_IsInGroup(c *gin.Context) {
	id := c.Query("id")
	group := c.Query("group")

	if id == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, api.GroupStatus{Ok: false, Message: "can't get id param"})
		return
	}

	if group == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, api.GroupStatus{Ok: false, Message: "can't get id param"})
		return
	}

	intId, err := strconv.Atoi(id)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, api.GroupStatus{Ok: false, Message: "bad id"})
		return
	}

	groupId, err := strconv.Atoi(group)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, api.GroupStatus{Ok: false, Message: "bad group id"})
		return
	}

	status, err := groups.Client.IsInGroup(context.Background(), &api.PropReq{
		GroupId: uint32(groupId),
		PropIds: []uint32{uint32(intId)},
	})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, api.GroupStatus{Ok: false, Message: "internal error"})
		return
	}

	c.JSON(http.StatusOK, status)
}

func Groups_AddPropsToGroup(c *gin.Context) {
	username, err := c.Cookie("username")

	if err != nil || username == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, api.GroupStatus{Ok: false, Message: "Can't get username from cookie"})
		return
	}

	prop := PropGroup{}

	rawData, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, api.GroupStatus{Ok: false, Message: "bad request body"})
		return
	}

	err = json.Unmarshal(rawData, &prop)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, api.GroupStatus{Ok: false, Message: "bad json"})
		return
	}

	gStatus, err := groups.Client.AddPropToGroup(context.Background(), &api.PropReq{
		GroupId: prop.GroupId,
		PropIds: prop.Ids,
	})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, api.GroupStatus{Ok: false, Message: "internal error"})
		return
	}

	props := &api2.ActionRequest{
		Note:     prop.Note,
		User:     username,
		SearchId: prop.GroupId,
	}

	for _, id := range prop.Ids {
		props.Properties = append(props.Properties, &api2.HistoryProperty{Id: id})
	}

	hStatus, err := history.Client.AddToGroup(context.Background(), props)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, api.GroupStatus{Ok: false, Message: "internal error"})
		return
	}

	if !gStatus.Ok || !hStatus.Ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, api.GroupStatus{Ok: false, Message: "internal error"})
		return
	}

	c.JSON(http.StatusOK, hStatus)
}

func Groups_RemoveFromGroup(c *gin.Context) {
	username, err := c.Cookie("username")

	if err != nil || username == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			api.GroupStatus{Ok: false, Message: "Can't get username from cookie"})
		return
	}

	prop := PropGroup{}

	rawData, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, api.GroupStatus{Ok: false, Message: "bad request body"})
		return
	}

	err = json.Unmarshal(rawData, &prop)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, api.GroupStatus{Ok: false, Message: "bad json"})
		return
	}

	gStatus, err := groups.Client.RemovePropFromGroup(context.Background(), &api.PropReq{
		GroupId: prop.GroupId,
		PropIds: prop.Ids,
	})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, api.GroupStatus{Ok: false, Message: "internal error"})
		return
	}

	props := &api2.ActionRequest{
		Note:     prop.Note,
		User:     username,
		SearchId: prop.GroupId,
	}

	for _, id := range prop.Ids {
		props.Properties = append(props.Properties, &api2.HistoryProperty{Id: id})
	}

	hStatus, err := history.Client.RemoveFromGroup(context.Background(), props)

	if !gStatus.Ok || !hStatus.Ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, api.GroupStatus{Ok: false, Message: "internal error"})
		return
	}

	c.JSON(http.StatusOK, hStatus)
}
