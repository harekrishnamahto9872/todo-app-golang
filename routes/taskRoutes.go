package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/harekrishnamahto9872/todo-app-golang/controllers/task"
	"github.com/harekrishnamahto9872/todo-app-golang/controllers/user"
	"github.com/harekrishnamahto9872/todo-app-golang/middleware"
	"go.mongodb.org/mongo-driver/mongo"
)

// SetRoutes //
// creates gin router and defines routes
func SetRoutes(router *gin.Engine, client *mongo.Client) {
	tasks := router.Group("/api/v1/tasks")

	tasks.GET("", middleware.Authentication(), func(c *gin.Context) {
		task.GetAllTasks(c, client)
	})
	tasks.POST("", middleware.Authentication(), func(c *gin.Context) {
		task.CreateTask(c, client)
	})
	tasks.DELETE("/:id", middleware.Authentication(), func(c *gin.Context) {
		task.DeleteTask(c, client)
	})
	tasks.PUT("/:id", middleware.Authentication(), func(c *gin.Context) {
		task.UpdateTask(c, client)
	})

	//auth routes
	authentication := router.Group("/api/v1/auth")

	authentication.POST("/login", func(c *gin.Context) {
		user.Login(c, client)
	})
	authentication.POST("/register", func(c *gin.Context) {
		user.Register(c, client)
	})
	authentication.GET("/logout", user.Logout)

}
