package controllers

import (
	"crud-api/models"
	"crud-api/utils/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CurrentUser(c *gin.Context) {

	// gets user id from extract token id function
	user_id, err := token.ExtractTokenID(c)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// get user from database by their id
	u, err := models.GetUserByID(user_id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// pass user with http response
	c.JSON(http.StatusOK, gin.H{"message": "success", "data": u})
}
