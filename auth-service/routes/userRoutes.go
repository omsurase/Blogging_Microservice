package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/omsurase/Blogging_Microservice/auth-service/controllers"
	"github.com/omsurase/Blogging_Microservice/auth-service/middleware"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate)
	incomingRoutes.GET("users/", controller.GetUsers())
	incomingRoutes.GET("users/:user_id", controller.GetUser())
}
