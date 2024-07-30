package controllers

import (
	"net/http"

	"github.com/ashcoder666/konnect/services"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateWallet(c *gin.Context) {

	type walletInput struct {
		user_id uuid.UUID
	}

	var walletInstance walletInput

	if err := c.BindJSON(&walletInstance); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Bad Request Body"})
		return
	}

	err := services.CreateWallet(walletInstance.user_id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Bad Request Body"})
		return
	}

}
