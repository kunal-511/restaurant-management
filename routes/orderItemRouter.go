package routes

import (
	"go-restaurant-management/controllers"

	"github.com/gin-gonic/gin"
)

func OrderItemRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/order_items", controllers.GetOrderItems())
	incomingRoutes.GET("/order_items/:order_item_id", controllers.GetOrderItem())
	incomingRoutes.GET("/order_items/order/:order_id", controllers.GetOrderItemsByOrder())
	incomingRoutes.POST("/order_items", controllers.CreateOrderItem())
	incomingRoutes.PATCH("/order_items/:order_item_id", controllers.UpdateOrderItem())
	incomingRoutes.DELETE("/order_items/:order_item_id", controllers.DeleteOrderItem())
}
