package controllers

import (
	"crud-api/models"
	"crud-api/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

// login function that returns http request and response
func Login(c *gin.Context) {

	// create needed fields to input variable
	var input types.LoginInput

	// if input is not entered
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{}

	user.Email = input.Email
	user.Password = input.Password

	// loginc check method checks email and password and returns a token
	token, err := models.LoginCheck(user.Email, user.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email or password is incorrect."})
		return
	}

	// if login wa successful we pass the token for the user to get authenticated
	c.JSON(http.StatusOK, gin.H{"token": token})
}

// registiring user
func Register(c *gin.Context) {

	var input types.RegisterInput

	// checking if all fields are filled
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{}
	/*
		Here I'm assging values of input to the models struct and calling saveuser fucntion that will hash the password
		and save user to the database
	*/
	user.FirstName = input.First_name
	user.LastName = input.Last_name
	user.Email = input.Email
	user.Password = input.Password

	// calling save functions
	_, err := user.SaveUser()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 200 ok
	c.JSON(http.StatusOK, gin.H{"message": "registered"})

}
