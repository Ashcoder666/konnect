package controllers

import (
	"github.com/ashcoder666/konnect/services"
	"github.com/gin-gonic/gin"
)

func ListAllPartners(c *gin.Context) {
	services.ListAllPartners()
}
