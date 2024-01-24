package controllers

import (
	"net/http"
	"rest-go-demo/database"
	"rest-go-demo/entity"
	"github.com/gin-gonic/gin"
)

// GetAllPerson gets all person data
func GetAllPerson(c *gin.Context) {
	var persons []entity.Person
	database.Connector.Find(&persons)
	c.JSON(http.StatusOK, persons)
}

// GetPersonByID returns person with specific ID
func GetPersonByID(c *gin.Context) {
	id := c.Param("id")

	var person entity.Person
	database.Connector.First(&person, id)
	c.JSON(http.StatusOK, person)
}

// CreatePerson creates person
func CreatePerson(c *gin.Context) {
	var person entity.Person
	if err := c.ShouldBindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.Connector.Create(&person)
	c.JSON(http.StatusCreated, person)
}

// UpdatePersonByID updates person with respective ID
func UpdatePersonByID(c *gin.Context) {
	var person entity.Person
	if err := c.ShouldBindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.Connector.Save(&person)
	c.JSON(http.StatusOK, person)
}

// DeletePersonByID deletes person with specific ID
func DeletePersonByID(c *gin.Context) {
	id := c.Param("id")

	var person entity.Person
	if err := database.Connector.Where("id = ?", id).Delete(&person).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}

	c.Status(http.StatusNoContent)
}
