package controllers

import (
	"crud-api/models"
	"crud-api/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostManagers(c *gin.Context) {
	//query := `INSERT INTO managers (first_name, last_name, phone_number, email) VALUES (?,?,?,?)`
	manager := models.Managers{}

	var input types.ManagerInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	manager.FirstName = input.First_name
	// later will be used to check errors
	if manager.FirstName == "test7@gmail.com" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "no"})
		return
	}
	manager.LastName = input.Last_name
	manager.PhoneNumber = input.Phone_number
	manager.Email = input.Email
	models.DB.Create(&manager)
	//models.DB.Raw(query, manager.FirstName, manager.LastName, manager.PhoneNumber, manager.Email)
	c.JSON(http.StatusOK, gin.H{"message": "200 ok"})
}
