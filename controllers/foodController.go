package controllers

import "github.com/gin-gonic/gin"

func GetFoods() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Get all foods",
		})
	}
}

func GetFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Get a food",
		})
	}
}

func CreateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Create a food",
		})
	}
}

func UpdateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Update a food",
		})
	}
}
func DeleteFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Delete a food",
		})
	}
}

func round(x float64) int {
	return int(x + 0.5)
}
func toFixed(x float64, precision int) float64 {
	return float64(int(x*100)) / 100
}
