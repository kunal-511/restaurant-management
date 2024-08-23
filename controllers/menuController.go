package controllers

import "github.com/gin-gonic/gin"

func GetMenus() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Get all menus",
		})
	}

}

func GetMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Get a menu",
		})
	}
}

func CreateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Create a menu",
		})
	}
}

func UpdateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Update a menu",
		})
	}
}
