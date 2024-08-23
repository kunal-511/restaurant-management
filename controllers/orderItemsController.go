package controllers

import "github.com/gin-gonic/gin"

func GetOrderItems() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Get all orderItems",
		})
	}

}

func GetOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Get a orderItem",
		})
	}
}

func CreateOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Create a orderItem",
		})
	}
}
func GetOrderItemsByOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Get all orderItems by order",
		})
	}
}

func UpdateOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Update a orderItem",
		})
	}
}

func DeleteOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Delete a orderItem",
		})
	}
}
