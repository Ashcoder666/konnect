package controllers

import (
	"fmt"
	"net/http"

	"github.com/ashcoder666/konnect/models"
	"github.com/ashcoder666/konnect/services"
	"github.com/ashcoder666/konnect/utils"
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

func UserLogin(c *gin.Context) {
	type loginInput struct {
		EmailorUsername string `json:"emailorusername"`
		Password        string `json:"password"`
	}

	var loginData loginInput

	if err := c.BindJSON(&loginData); err != nil {
		fmt.Println(loginData)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	fmt.Println(loginData)

	selectedUser, err := services.CheckUserExists(loginData.EmailorUsername)

	fmt.Println(selectedUser)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user doesnt exist"})
		return
	}

	if err := services.ComparePassword(loginData.Password, selectedUser.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid password"})
		return
	}

	token, err := utils.GenerateToken(selectedUser.Email)

	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "error in creating token"})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"message": "success", "Token": token})

}
