package routes

import (
	"auth-with-jwt/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(c *gin.Engine) {
	c.POST("users/signup", controllers.Signup())
	c.POST("users/login", controllers.Login())
}
