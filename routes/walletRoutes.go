package routes

import (
	"github.com/ashcoder666/konnect/controllers"
	"github.com/gin-gonic/gin"
)

func WalletRoutes(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		v1.POST("/user/wallet", controllers.CreateWallet)
	}
}
