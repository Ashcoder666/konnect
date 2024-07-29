package routes

import (
	"github.com/ashcoder666/konnect/controllers"
	"github.com/ashcoder666/konnect/middlewares"
	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		v1.GET("/partners/list", middlewares.AuthMiddleware(), controllers.ListAllPartners)
	}
}
