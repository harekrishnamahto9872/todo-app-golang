package task

import (
	"github.com/gin-gonic/gin"
	"github.com/harekrishnamahto9872/todo-app-golang/models"
	"github.com/harekrishnamahto9872/todo-app-golang/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetAllTasks //
// @desc Get all tasks
// @route GET /api/v1/tasks
// @access Private
func GetAllTasks(c *gin.Context, client *mongo.Client) {

	tasks := []models.Task{}

	tasksCollection := client.Database("Todo").Collection("tasks")

	// query db and filter by user id
	cursor, findErr := tasksCollection.Find(c.Request.Context(), bson.M{
		"user": c.Keys["id"],
	})
	if findErr != nil {
		c.JSON(500, util.ResError{
			Success: false,
			Error:   findErr,
		})
		return
	}

	// loop through cursor and put tasks in the tasks slice of tasks
	cursorErr := cursor.All(c.Request.Context(), &tasks)
	if cursorErr != nil {
		c.JSON(500, util.ResError{
			Success: false,
			Error:   cursorErr,
		})
		return
	}

	c.JSON(200, util.ResTasks{
		Success: true,
		Message: tasks,
	})
}
