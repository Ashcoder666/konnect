package controllers

import (
	"fmt"
	"net/http"

	"github.com/ashcoder666/konnect/models"
	"github.com/ashcoder666/konnect/services"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func UserRegistration(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		fmt.Println(user)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	if user.Email == "" || user.Password == "" || user.Role == "" || user.UserName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "All fields need to be send"})
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	user.Password = string(hashedPassword)

	fmt.Println(user)

	err := services.CreateUserService(&user)

	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "user created successfully"})

}
