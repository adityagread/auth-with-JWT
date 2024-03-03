package routes

import (
	"auth-with-jwt/controllers"
	"auth-with-jwt/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(c *gin.Engine) {
	c.Use(middleware.Authenticate())
	c.GET("/users", controllers.GetUsers())
	c.GET("/users/:user_id", controllers.GetUser())
}
