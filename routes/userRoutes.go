package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/harekrishnamahto9872/todo-app-golang/controllers/user"
	"go.mongodb.org/mongo-driver/mongo"
)

//defines user routes
func SetUserRoutes(router *gin.Engine, client *mongo.Client) {
	// routes for authorization and authentication
	authentication := router.Group("/api/v1/auth")

	authentication.POST("/login", func(c *gin.Context) {
		user.Login(c, client)
	})
	authentication.POST("/register", func(c *gin.Context) {
		user.Register(c, client)
	})
	authentication.GET("/logout", user.Logout)
}
