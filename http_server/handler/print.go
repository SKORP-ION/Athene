package handler

import (
	"Athena/http_server/gRPC/staff"
	"Athena/staff/api"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

func View_Print(c *gin.Context) {
	//data := PrintReq{}

	ticket := c.PostForm("ticket")

	if ticket == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	name := c.PostForm("name")

	if name == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	strRec := c.PostForm("recordId")

	if strRec == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	rec, err := strconv.Atoi(strRec)

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	resp, err := staff.Client.GetRecord(context.Background(), &api.SRecord{Id: uint32(rec)})

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	prop := Property{
		Inventory: resp.Prop.Inventory,
		Name:      name,
		Serial:    resp.Prop.Serial,
	}

	nameList := strings.Split(resp.Emp.Name, " ")

	emp := Employee{
		Table:      resp.Emp.Table,
		Job:        resp.Emp.Job,
		Department: resp.Emp.Department,
		Name:       resp.Emp.Name,
		ShortName: fmt.Sprintf("%s %s.%s.",
			nameList[0], string([]rune(nameList[1])[0]), string([]rune(nameList[2])[0])),
	}

	d := resp.Date.AsTime()

	c.HTML(http.StatusOK, "voucher.html", gin.H{
		"Prop":   prop,
		"date":   fmt.Sprintf("%d.%d.%d", d.Day(), d.Month(), d.Year()),
		"Emp":    emp,
		"Ticket": ticket,
	})
}
