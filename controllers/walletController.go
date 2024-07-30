package controllers

import (
	"fmt"
	"net/http"

	"github.com/ashcoder666/konnect/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateWallet(c *gin.Context) {

	type walletInput struct {
		UserId uuid.UUID `json:"user_id"`
	}

	var walletInstance walletInput

	if err := c.BindJSON(&walletInstance); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Bad Request Body"})
		return
	}

	fmt.Println(walletInstance)

	err := services.CreateWallet(walletInstance.UserId)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Bad Request Body"})
		return
	}

}
