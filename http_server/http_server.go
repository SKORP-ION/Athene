package main

import (
	"Athena/http_server/handler"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	gin.SetMode(gin.DebugMode)
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.Static("/static", "static")
	router.StaticFile("/favicon.ico", "favicon.ico")

	router.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})
	router.POST("/login", handler.Authentication)

	view := router.Group("/view")
	view.Use(handler.AuthorizationMiddleware)
	{
		view.GET("/", func(c *gin.Context) {
			c.Redirect(http.StatusMovedPermanently, "/view/getProperty")
		})
		view.GET("/r12Import", func(c *gin.Context) {
			c.HTML(http.StatusOK, "newPropertyGet.html", nil)
		})
		view.GET("/getProperty", func(c *gin.Context) {
			c.HTML(http.StatusOK, "getProperty.html", nil)
		})
		view.GET("/giveToEmployee", func(c *gin.Context) {
			c.HTML(http.StatusOK, "giveToEmployee.html", nil)
		})
		view.GET("/property", func(c *gin.Context) {
			c.HTML(http.StatusOK, "property.html", nil)
		})
		view.GET("/action", handler.View_Action)
		view.GET("/groups", func(c *gin.Context) {
			c.HTML(http.StatusOK, "groups.html", nil)
		})
		view.GET("/group", func(c *gin.Context) {
			c.HTML(http.StatusOK, "group.html", nil)
		})
		view.GET("/staff", func(c *gin.Context) {
			c.HTML(http.StatusOK, "staff.html", nil)
		})
		view.POST("/print", handler.View_Print)
	}
	router.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusTemporaryRedirect, "/view/")
	})

	private := router.Group("/private")
	private.Use(handler.AuthorizationMiddleware)
	{
		private.GET("/logout", handler.Logout)

		private.GET("/property/getWarehouses", handler.Property_GetWarehouses)
		private.GET("/property/getActions", handler.Property_GetActions)
		private.POST("/property/getCount", handler.Property_GetCount)
		private.POST("/property/getProperty", handler.Property_GetProperties)
		private.GET("/property/getOneProperty", handler.Property_GetOneProperty)
		private.GET("/property/isOnWarehouse", handler.Property_IsOnWarehouse)
		private.POST("/property/sendToWarehouse", handler.Property_SendToWarehouse)

		private.POST("/ldap/getStaff", handler.SearchStaff)

		private.POST("/history/CreateCard", handler.History_CreateCard)
		private.POST("/history/InstallOnWorkspace", handler.History_DoAction)
		private.POST("/history/RemoveFromWorkspace", handler.History_DoAction)
		private.POST("/history/NeedRepair", handler.History_DoAction)
		private.POST("/history/SendToRepair", handler.History_DoAction)
		private.POST("/history/ReceiveFromRepair", handler.History_DoAction)
		private.POST("/history/Archive", handler.History_DoAction)
		private.POST("/history/DeArchive", handler.History_DoAction)
		private.POST("/history/ChangeName", handler.History_DoAction)
		private.POST("/history/ChangeInventory", handler.History_DoAction)
		private.GET("/history/getHistory", handler.History_GetHistory)
		private.GET("/history/isCreated", handler.History_IsCreated)
		private.GET("/history/isOnWorkspace", handler.History_IsOnWorkspace)
		private.GET("/history/isNeedsRepair", handler.History_IsNeedsRepair)
		private.GET("/history/isUnderRepair", handler.History_IsUnderRepair)
		private.GET("/history/isInStock", handler.History_IsInStock)
		private.GET("/history/isInArchive", handler.History_IsInArchive)

		private.GET("/groups/GetGroups", handler.Groups_GetGroups)
		private.POST("/groups/CreateGroup", handler.Groups_GroupAction)
		private.POST("/groups/RemoveGroup", handler.Groups_GroupAction)
		private.GET("/groups/IsInGroup", handler.Groups_IsInGroup)
		private.POST("/groups/AddPropsToGroup", handler.Groups_AddPropsToGroup)
		private.POST("/groups/RemoveFromGroup", handler.Groups_RemoveFromGroup)

		private.POST("/staff/GetStaff", handler.Staff_GetStaff)
		private.GET("/staff/GetStaffProps", handler.Staff_GetStaffProps)
		private.POST("/staff/GetCount", handler.Staff_GetCount)
		private.GET("/staff/isWithEmployee", handler.Staff_IsWithEmployee)
		private.POST("/staff/GiveToEmployee", handler.Staff_GiveToEmployee)
		private.POST("/staff/TakeFromEmployee", handler.Staff_TakeFromEmployee)
		private.POST("/staff/CreateEmployee", handler.Staff_CreateEmployee)
	}

	err := router.RunTLS(handler.Host, handler.Cert, handler.PrivKey)

	if err != nil {
		panic(err)
	}
}
