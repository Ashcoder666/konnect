package routes

import (
	"github.com/ashcoder666/konnect/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		v1.POST("/signup", controllers.UserRegistration)
		v1.POST("/login", controllers.UserLogin)
		v1.POST("/forgot-password", controllers.UserForgotPassword)
		v1.POST("/reset-password", controllers.UserResetPassword)
	}
}
