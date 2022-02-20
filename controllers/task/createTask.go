package task

import (
	"github.com/gin-gonic/gin"
	"github.com/harekrishnamahto9872/todo-app-golang/models"
	"github.com/harekrishnamahto9872/todo-app-golang/util"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// CreateTask //
// @desc Create a new task from task struct and add it to the database
// @route POST /api/v1/tasks
// @access Private
func CreateTask(c *gin.Context, client *mongo.Client) {

	newTask := models.NewTask{}
	newTask.SetCreatedAt()
	newTask.SetUpdatedAt()
	newTask.User = c.Keys["id"].(string)

	bindErr := c.ShouldBindJSON(&newTask)
	if bindErr != nil {
		c.JSON(400, util.ResError{
			Success: false,
			Error:   bindErr,
		})
		return
	}

	taskCollection := client.Database("Todo").Collection("tasks")

	// put task in the collection, returns _id
	taskRes, insertErr := taskCollection.InsertOne(c.Request.Context(), newTask)
	if insertErr != nil {
		c.JSON(401, util.ResError{
			Success: false,
			Error:   insertErr,
		})
		return
	}

	// uses struct to format the response body. ID is from InsertOne
	task := models.Task{
		ID:          taskRes.InsertedID.(primitive.ObjectID),
		Title:       newTask.Title,
		Description: newTask.Description,
		DueDate:     newTask.DueDate,
		CreatedAt:   newTask.CreatedAt,
		UpdatedAt:   newTask.UpdatedAt,
		User:        newTask.User,
	}

	c.JSON(200, util.ResTask{
		Success: true,
		Message: task,
	})
}
