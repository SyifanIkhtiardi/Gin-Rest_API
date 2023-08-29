package main

import (
	"golang-projects/web-service-gin/controlers"
	"golang-projects/web-service-gin/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/employee", controlers.AllEmployee)
	r.GET("/employee/:id", controlers.Employee)
	r.POST("/employee", controlers.CreateNewEmployee)
	r.PUT("/employee/:id", controlers.UpdateEmployee)
	r.DELETE("/employee/:id", controlers.DeleteEmployee)

	r.Run()
}
