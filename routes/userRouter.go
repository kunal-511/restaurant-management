package routes

import (
	"go-restaurant-management/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingROutes *gin.Engine) {

	incomingROutes.GET("/users", controllers.GetUsers())
	incomingROutes.GET("/users/:user_id", controllers.GetUser())
	incomingROutes.POST("/users/signup", controllers.CreateUser())
	incomingROutes.POST("/users/login", controllers.LoginUser())
}
