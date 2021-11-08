package handler

import (
	"Athena/history/api"
	"Athena/http_server/gRPC/history"
	"context"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
)

func History_CreateCard(c *gin.Context) {
	status := &api.Status{}
	username, err := c.Cookie("username")

	if err != nil {
		c.Error(err)
		status.Ok = false
		status.Message = "Can't read username from cookie"
		c.JSON(http.StatusBadRequest, &status)
		return
	}

	rawData, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		c.Error(err)
		status.Ok = false
		status.Message = "Empty request body"
		c.JSON(http.StatusBadRequest, &status)
		return
	}

	data := AddPropertyRequest{}

	err = json.Unmarshal(rawData, &data)

	if err != nil {
		c.Error(err)
		status.Ok = false
		status.Message = "Bad json"
		c.JSON(http.StatusBadRequest, &status)
		return
	}

	requestData := &api.ActionRequest{}

	requestData.User = username
	requestData.Note = data.Note

	for _, propData := range data.Properties {
		property := &api.HistoryProperty{
			Inventory: propData.Inventory,
			Serial:    propData.Serial,
			Name:      propData.Name,
		}
		requestData.Properties = append(requestData.Properties, property)
	}
	statusWithProp, err := history.Client.CreateCard(context.Background(), requestData)

	if err != nil {
		status.Ok = false
		status.Message = "internal error"
		c.JSON(http.StatusInternalServerError, &status)
		return
	}

	c.JSON(http.StatusOK, statusWithProp)
}

func History_DoAction(c *gin.Context) {
	status := &api.Status{}
	username, err := c.Cookie("username")

	if err != nil {
		c.Error(err)
		status.Ok = false
		status.Message = "Can't read username from cookie"
		c.JSON(http.StatusBadRequest, &status)
		return
	}

	rawData, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		c.Error(err)
		status.Ok = false
		status.Message = "Empty request body"
		c.JSON(http.StatusBadRequest, &status)
		return
	}

	data := AddPropertyRequest{}

	err = json.Unmarshal(rawData, &data)

	if err != nil {
		c.Error(err)
		status.Ok = false
		status.Message = "Bad json"
		c.JSON(http.StatusBadRequest, &status)
		return
	}

	requestData := &api.ActionRequest{}

	requestData.User = username
	requestData.Note = data.Note

	for _, propData := range data.Properties {
		property := &api.HistoryProperty{
			Id:        propData.Id,
			Inventory: propData.Inventory,
			Serial:    propData.Serial,
			Name:      propData.Name,
		}
		requestData.Properties = append(requestData.Properties, property)
	}

	path := c.FullPath()

	status = &api.Status{}

	switch path {
	case "/private/history/InstallOnWorkspace":
		status, err = history.Client.InstallOnWorkspace(context.Background(), requestData)
	case "/private/history/RemoveFromWorkspace":
		status, err = history.Client.RemoveFromWorkspace(context.Background(), requestData)
	case "/private/history/NeedRepair":
		status, err = history.Client.NeedRepair(context.Background(), requestData)
	case "/private/history/SendToRepair":
		status, err = history.Client.SendToRepair(context.Background(), requestData)
	case "/private/history/ReceiveFromRepair":
		status, err = history.Client.ReceiveFromRepair(context.Background(), requestData)
	case "/private/history/Archive":
		status, err = history.Client.Archive(context.Background(), requestData)
	case "/private/history/DeArchive":
		status, err = history.Client.DeArchive(context.Background(), requestData)
	case "/private/history/ChangeName":
		status, err = history.Client.ChangeName(context.Background(), requestData)
	case "/private/history/ChangeInventory":
		status, err = history.Client.ChangeInventory(context.Background(), requestData)
	}

	if err != nil {
		status.Ok = false
		status.Message = "Internal error"
		c.JSON(http.StatusInternalServerError, &status)
		return
	}

	c.JSON(http.StatusOK, status)
}

func History_GetHistory(c *gin.Context) {
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

	records, err := history.Client.GetHistory(context.Background(), &api.HistoryRequest{Id: uint32(id)})

	if err != nil {
		err := errors.New("internal error")
		c.Error(err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, records.Records)
}

func History_IsCreated(c *gin.Context) {
	serial := c.Query("serial")

	if serial == "" {
		err := errors.New("can't get serial param")
		c.Error(err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	status, err := history.Client.IsCreated(context.Background(), &api.PropertySerial{Serial: serial})

	if err != nil {
		err := errors.New("internal error")
		c.Error(err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, status)
}

func History_IsOnWorkspace(c *gin.Context) {
	serial := c.Query("serial")

	if serial == "" {
		err := errors.New("can't get serial param")
		c.Error(err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	status, err := history.Client.IsOnWorkspace(context.Background(), &api.PropertySerial{Serial: serial})

	if err != nil {
		err := errors.New("internal error")
		c.Error(err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, status)
}

func History_IsNeedsRepair(c *gin.Context) {
	serial := c.Query("serial")

	if serial == "" {
		err := errors.New("can't get serial param")
		c.Error(err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	status, err := history.Client.IsNeedsRepair(context.Background(), &api.PropertySerial{Serial: serial})

	if err != nil {
		err := errors.New("internal error")
		c.Error(err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, status)
}

func History_IsUnderRepair(c *gin.Context) {
	serial := c.Query("serial")

	if serial == "" {
		err := errors.New("can't get serial param")
		c.Error(err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	status, err := history.Client.IsUnderRepair(context.Background(), &api.PropertySerial{Serial: serial})

	if err != nil {
		err := errors.New("internal error")
		c.Error(err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, status)
}

func History_IsInArchive(c *gin.Context) {
	serial := c.Query("serial")

	if serial == "" {
		err := errors.New("can't get serial param")
		c.Error(err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	status, err := history.Client.IsInArchive(context.Background(), &api.PropertySerial{Serial: serial})

	if err != nil {
		err := errors.New("internal error")
		c.Error(err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, status)
}

func History_IsInStock(c *gin.Context) {
	serial := c.Query("serial")

	if serial == "" {
		err := errors.New("can't get serial param")
		c.Error(err)
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	status, err := history.Client.IsInStock(context.Background(), &api.PropertySerial{Serial: serial})

	if err != nil {
		err := errors.New("internal error")
		c.Error(err)
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, status)
}
