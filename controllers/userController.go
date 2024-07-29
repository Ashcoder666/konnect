package controllers

import (
	"net/http"

	"github.com/ashcoder666/konnect/services"
	"github.com/gin-gonic/gin"
)

func ListAllPartners(c *gin.Context) {
	allPartners, err := services.ListAllPartners()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err})
	}

	c.JSON(http.StatusAccepted, gin.H{"Message": "success", "Data": allPartners})
}
