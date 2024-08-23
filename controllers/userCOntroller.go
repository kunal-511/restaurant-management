package controllers

import "github.com/gin-gonic/gin"

func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Get all users",
		})
	}

}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Get a user",
		})
	}
}

func CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Create a user",
		})
	}
}
func LoginUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Login a user",
		})
	}
}

func HashPassword(password string) string {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hash a password",
		})
	}
}

func ComparePassword(userPassword string, providerPasswrod string) (bool, )string {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Compare a password",
		})
	}
}
