package routes

import (
	controller "golang/battery-tracking/controller"

	"github.com/gin-gonic/gin"
	"golang/battery-tracking/middleware"
)

func UserRoutes (incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.AuthMiddleware())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/:user_id", controller.GetUser())
}