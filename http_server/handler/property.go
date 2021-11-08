package handler

import (
	api2 "Athena/history/api"
	"Athena/http_server/gRPC/history"
	"Athena/http_server/gRPC/property"
	"Athena/property/api"
	"context"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

func Property_GetProperties(c *gin.Context) {
	params, err := ReadSearchParams(c.Request.Body)

	if err != nil {
		c.Error(err)
	}

	properties, err := property.Client.GetProperty(context.Background(), params)

	if err != nil {
		c.Error(err)
	}

	c.JSON(http.StatusOK, properties)
}

func Property_GetWarehouses(c *gin.Context) {
	warehouses, err := property.Client.GetWarehouses(context.Background(), &api.EmptyPropertyRequest{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{Data: "Internal error."})
		return
	}

	//response, err := json.Marshal(warehouses.Data)
	c.JSON(http.StatusOK, warehouses.Data)
}

func Property_GetCount(c *gin.Context) {
	params, err := ReadSearchParams(c.Request.Body)

	if err != nil {
		c.Error(err)
	}

	params.Offset = 0
	params.Limit = 0

	count, err := property.Client.GetCount(context.Background(), params)

	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Data: "Can't get count of properties. Connect to your system administrator."})
		return
	}

	c.JSON(http.StatusOK, count)
}

func Property_GetActions(c *gin.Context) {
	actions, err := property.Client.GetActionsList(context.Background(), &api.EmptyPropertyRequest{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{Data: "Internal error."})
		return
	}

	c.JSON(http.StatusOK, actions.Actions)
}

func Property_GetOneProperty(c *gin.Context) {
	strId := c.Query("id")

	if strId == "" {
		err := errors.New("can't get id param")
		c.Error(err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	id, err := strconv.Atoi(strId)

	if err != nil {
		err := errors.New("can't convert id to int")
		c.Error(err)
		c.AbortWithError(http.StatusBadRequest, errors.New("can't convert id to int"))
		return
	}

	prop, err := property.Client.GetOneProperty(context.Background(), &api.GetOnePropertyRequest{Id: uint32(id)})

	if err != nil {
		c.Error(errors.New("internal error"))
		c.AbortWithError(http.StatusInternalServerError, errors.New("internal server error"))
		return
	}
	c.JSON(http.StatusOK, prop)
}

func Property_IsOnWarehouse(c *gin.Context) {
	strId := c.Query("id")
	strWarehouse := c.Query("warehouse")

	if strId == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, api.PropStatus{Ok: false, Message: "can't get id"})
		return
	}

	if strWarehouse == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, api.PropStatus{Ok: false, Message: "can't get warehouse id"})
		return
	}

	id, err := strconv.Atoi(strId)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, api.PropStatus{Ok: false, Message: "bad id"})
		return
	}

	warehouse, err := strconv.Atoi(strWarehouse)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, api.PropStatus{Ok: false, Message: "bad warehouse id"})
		return
	}

	status, err := property.Client.IsOnWarehouse(context.Background(), &api.IsInWarehouseReq{
		WarehouseId: uint32(warehouse),
		PropertyId:  uint32(id),
	})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, api.PropStatus{Ok: false, Message: "internal error"})
		return
	}

	c.JSON(http.StatusOK, status)
}

func Property_SendToWarehouse(c *gin.Context) {
	username, err := c.Cookie("username")

	if err != nil || username == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, api.PropStatus{Ok: false, Message: "Can't get username from cookie"})
		return
	}

	prop := SendToWarehouseReq{}

	rawData, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, api.PropStatus{Ok: false, Message: "bad request body"})
		return
	}

	err = json.Unmarshal(rawData, &prop)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, api.PropStatus{Ok: false, Message: "bad json"})
		return
	}

	pStatus, err := property.Client.SendToWarehouse(context.Background(), &api.SendToWarhouseReq{
		PropertiesId: prop.Ids,
		WarehouseId:  prop.WarehouseId,
	})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, api.PropStatus{Ok: false, Message: "internal error"})
		return
	}

	properties := make([]*api2.HistoryProperty, 0)

	for _, id := range prop.Ids {
		properties = append(properties, &api2.HistoryProperty{Id: id})
	}

	hStatus, err := history.Client.SendToWarehouse(context.Background(), &api2.ActionRequest{
		Note:       prop.Note,
		Properties: properties,
		User:       username,
		SearchId:   prop.WarehouseId,
	})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, api.PropStatus{Ok: false, Message: "internal error"})
		return
	}

	if !pStatus.Ok || !hStatus.Ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, api.PropStatus{Ok: false, Message: "internal error"})
		return
	}

	c.JSON(http.StatusOK, hStatus)
}

func View_Action(c *gin.Context) {
	c.HTML(http.StatusOK, "actionProperty.html", nil)
}

func ReadSearchParams(body io.Reader) (*api.SearchParams, error) {
	rawData, err := ioutil.ReadAll(body)

	if err != nil {
		return nil, err
	}

	params := api.SearchParams{}

	err = json.Unmarshal(rawData, &params)

	if err != nil {
		return nil, err
	}

	return &params, nil
}
