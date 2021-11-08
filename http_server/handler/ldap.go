package handler

import (
	"Athena/http_server/gRPC/ldap"
	"Athena/http_server/gRPC/staff"
	"Athena/ldap/api"
	api2 "Athena/staff/api"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func SearchStaff(c *gin.Context) {
	rawData, err := ioutil.ReadAll(c.Request.Body)

	if err != nil {
		c.JSON(http.StatusBadRequest, Response{Data: "Empty request body"})
		return
	}

	data := make(map[string]interface{})

	err = json.Unmarshal(rawData, &data)

	if err != nil {
		c.JSON(http.StatusBadRequest, Response{Data: "Wrong json. Can't parse"})
		return
	}

	name := data["name"].(string)

	ldapStaff, err := ldap.Client.SearchEmployee(context.Background(), &api.SearchRequest{Name: name})

	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{Data: "Internal error."})
		return
	}

	dbStaff, err := staff.Client.GetStaff(context.Background(), &api2.GetReq{Search: name})

	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{Data: "Internal error."})
		return
	}

	response := make([]*api2.StaffEmployee, 0)

	response = dbStaff.Staff

	for _, emp := range ldapStaff.Staff {
		if !Include(response, emp) {
			response = append(response, &api2.StaffEmployee{
				Table:      emp.EmployeeNumber,
				Name:       fmt.Sprintf("*%s", emp.Name),
				Manager:    emp.Manager,
				Department: emp.Department,
				Job:        emp.JobTitle,
			})
		}
	}

	c.JSON(http.StatusOK, response)
}

func Include(array []*api2.StaffEmployee, target *api.LdapEmployee) bool {
	for _, emp := range array {
		if emp.Table == target.EmployeeNumber {
			return true
		}
	}
	return false
}
