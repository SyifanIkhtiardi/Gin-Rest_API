package controlers

import (
	"golang-projects/web-service-gin/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler function that return all employee in database
func AllEmployee(c *gin.Context) {
	var employee []models.Employee
	models.DB.Find(&employee)

	c.JSON(http.StatusOK, gin.H{"data": employee})
}

// Handler function to get employee by id
func Employee(c *gin.Context) {
	var employee models.Employee

	if err := models.DB.Where("id = ?", c.Param("id")).First(&employee).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Employee not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": employee})
}

// Handler function to create new employee
func CreateNewEmployee(c *gin.Context) {
	// Validate input
	var input models.CreateEmployeeInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create New Employee
	employee := models.Employee{Name: input.Name, Role: input.Role}
	models.DB.Create(&employee)

	c.JSON(http.StatusCreated, gin.H{"data": employee, "status": "Create new employee success"})
}

// Handler function to update employee data
func UpdateEmployee(c *gin.Context) {
	// Get model if exist
	var employee models.Employee
	if err := models.DB.Where("id = ?", c.Param("id")).First(&employee).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	// Validate input
	var input models.UpdateEmployeeInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&employee).Updates(input)

	c.JSON(http.StatusOK, gin.H{"status": "Successfully update employee data", "data": employee})
}

// Handler function to delete employee data
func DeleteEmployee(c *gin.Context) {
	// Get model if exist
	var employee models.Employee
	if err := models.DB.Where("id = ?", c.Param("id")).First(&employee).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found"})
		return
	}

	models.DB.Delete(&employee)

	c.JSON(http.StatusOK, gin.H{"status": "Successfully delete employee data"})
}
