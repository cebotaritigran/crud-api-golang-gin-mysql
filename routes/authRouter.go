package routes

import (
	"crud-api/controllers"

	"github.com/gin-gonic/gin"
)

// open routes that anybody can access
// signup has to be protected later
func AuthRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("users/signup", controllers.Register)
	incomingRoutes.POST("users/login", controllers.Login)
}
