package handler

import (
	api2 "Athena/history/api"
	"Athena/http_server/gRPC/history"
	"Athena/http_server/gRPC/staff"
	"Athena/staff/api"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
)

func Staff_CreateEmployee(c *gin.Context) {
	emp := api.StaffEmployee{}

	rawData, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &api.StatusWithEmployee{
			Ok:       false,
			Employee: nil,
		})
		return
	}

	err = json.Unmarshal(rawData, &emp)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &api.StatusWithEmployee{
			Ok:       false,
			Employee: nil,
		})
		return
	}

	resp, err := staff.Client.CreateEmployee(context.Background(), &emp)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &api.StatusWithEmployee{
			Ok:       false,
			Employee: nil,
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func Staff_GetStaff(c *gin.Context) {
	search := &api.GetReq{}

	rawData, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &api.GetResp{Ok: false, Staff: nil})
		return
	}

	err = json.Unmarshal(rawData, &search)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &api.GetResp{Ok: false, Staff: nil})
		return
	}

	employees, err := staff.Client.GetStaff(context.Background(), search)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &api.GetResp{Ok: false, Staff: nil})
		return
	}

	c.JSON(http.StatusOK, employees)
}

func Staff_GetStaffProps(c *gin.Context) {
	strId := c.Query("id")

	if strId == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, &api.SStatus{
			Ok:      false,
			Message: "can't get id param",
		})
		return
	}

	id, err := strconv.Atoi(strId)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &api.SStatus{
			Ok:      false,
			Message: "bad id",
		})
		return
	}

	req := api.StaffEmployee{Id: uint32(id)}

	status, err := staff.Client.GetStaffProp(context.Background(), &req)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &api.SStatus{
			Ok:      false,
			Message: "internal error",
		})
		return
	}

	c.JSON(http.StatusOK, status)
}

func Staff_GetCount(c *gin.Context) {
	search := &api.GetReq{}

	rawData, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &api.GetResp{Ok: false, Staff: nil})
		return
	}

	err = json.Unmarshal(rawData, &search)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &api.GetResp{Ok: false, Staff: nil})
		return
	}

	resp, err := staff.Client.GetCount(context.Background(), search)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &api.GetResp{Ok: false, Staff: nil})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func Staff_GiveToEmployee(c *gin.Context) {
	username, err := c.Cookie("username")

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &api.SStatus{Ok: false, Message: "can't get username"})
		return
	}

	rawData, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &api.SStatus{Ok: false, Message: "bad request body"})
		return
	}

	data := GiveReq{}

	err = json.Unmarshal(rawData, &data)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &api.SStatus{Ok: false, Message: "bad json"})
		return
	}

	SReq := api.GiveReq{EmployeeId: data.EmployeeId, Username: username}

	for _, id := range data.Ids {
		SReq.Ids = append(SReq.Ids, id)
	}

	Sstatus := make(chan *api.SStatus)
	Serr := make(chan error)

	go func() {
		//defer close(Serr)
		//defer close (Sstatus)
		status, err := staff.Client.GiveToEmployee(context.Background(), &SReq)
		Serr <- err
		Sstatus <- status
	}()

	HReq := api2.ActionRequest{
		User:     username,
		Note:     data.Note,
		SearchId: data.EmployeeId,
	}

	for _, id := range data.Ids {
		HReq.Properties = append(HReq.Properties, &api2.HistoryProperty{Id: id})
	}

	Hstatus := make(chan *api2.Status)
	Herr := make(chan error)

	go func() {
		//defer close(Herr)
		//defer close (Hstatus)
		status, err := history.Client.GiveToEmployee(context.Background(), &HReq)
		Herr <- err
		Hstatus <- status
	}()

	resp := &api.SStatus{
		Ok:      true,
		Message: "Success",
	}

	for i := 0; i < 2; i++ {
		select {
		case err := <-Serr:
			{
				if err != nil {
					c.AbortWithStatusJSON(http.StatusInternalServerError, &api.SStatus{Ok: false, Message: "internal error"})
					return
				}
			}
		case err := <-Herr:
			{
				if err != nil {
					c.AbortWithStatusJSON(http.StatusInternalServerError, &api.SStatus{Ok: false, Message: "internal error"})
					return
				}
			}
		}

		select {
		case status := <-Sstatus:
			{
				if !status.Ok {
					c.AbortWithStatusJSON(http.StatusInternalServerError, &api.SStatus{Ok: false, Message: "internal error"})
					return
				}
			}
		case status := <-Hstatus:
			{
				if !status.Ok {
					c.AbortWithStatusJSON(http.StatusInternalServerError, &api.SStatus{Ok: false, Message: "internal error"})
					return
				}
			}
		}
	}

	c.JSON(http.StatusOK, resp)
}

func Staff_TakeFromEmployee(c *gin.Context) {
	username, err := c.Cookie("username")

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &api.SStatus{Ok: false, Message: "can't get username"})
		return
	}

	rawData, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &api.SStatus{Ok: false, Message: "bad request body"})
		return
	}

	data := GiveReq{}

	err = json.Unmarshal(rawData, &data)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &api.SStatus{Ok: false, Message: "bad json"})
		return
	}

	SReq := api.GiveReq{EmployeeId: data.EmployeeId, Username: username}

	for _, id := range data.Ids {
		SReq.Ids = append(SReq.Ids, id)
	}

	Sstatus, err := staff.Client.TakeFromEmployee(context.Background(), &SReq)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &api.SStatus{Ok: false, Message: "internal error"})
		return
	}

	HReq := api2.ActionRequest{
		User:     username,
		Note:     data.Note,
		SearchId: data.EmployeeId,
	}

	for _, id := range data.Ids {
		HReq.Properties = append(HReq.Properties, &api2.HistoryProperty{Id: id})
	}

	Hstatus, err := history.Client.TakeFromEmployee(context.Background(), &HReq)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &api.SStatus{Ok: false, Message: "internal error"})
		return
	}

	if !Sstatus.Ok || !Hstatus.Ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &api.SStatus{Ok: false, Message: "internal error"})
		return
	}

	c.JSON(http.StatusOK, Sstatus)
}

func Staff_IsWithEmployee(c *gin.Context) {
	strId := c.Query("id")

	if strId == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, &api.SStatus{
			Ok:      false,
			Message: "can't get id param",
		})
		return
	}

	strEmp := c.Query("employee")

	if strEmp == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, &api.SStatus{
			Ok:      false,
			Message: "can't get employee param",
		})
		return
	}

	id, err := strconv.Atoi(strId)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &api.SStatus{
			Ok:      false,
			Message: "bad id",
		})
		return
	}

	emp, err := strconv.Atoi(strEmp)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, &api.SStatus{
			Ok:      false,
			Message: "bad employee id",
		})
		return
	}

	SReq := api.GiveReq{
		EmployeeId: uint32(emp),
		Ids:        []uint32{uint32(id)},
	}

	status, err := staff.Client.IsWithEmployee(context.Background(), &SReq)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, &api.SStatus{
			Ok:      false,
			Message: "internal error",
		})
		return
	}

	c.JSON(http.StatusOK, status)
}
