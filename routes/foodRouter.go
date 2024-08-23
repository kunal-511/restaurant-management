package routes

import (
	"go-restaurant-management/controllers"

	"github.com/gin-gonic/gin"
)

func FoodRoutes(incomingROutes *gin.Engine) {

	incomingROutes.GET("/foods", controllers.GetFoods())
	incomingROutes.GET("/foods/:food_id", controllers.GetFood())
	incomingROutes.POST("/foods", controllers.CreateFood())
	incomingROutes.PATCH("/foods/:food_id", controllers.UpdateFood())
	incomingROutes.DELETE("/foods/:food_id", controllers.DeleteFood())
}
