package task

import (
	"github.com/harekrishnamahto9872/todo-app-golang/util"

	"github.com/harekrishnamahto9872/todo-app-golang/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

// DeleteTask //
// @desc Delete a task with the id in the request params
// @route DELETE /api/v1/tasks/:id
// @access Private
func DeleteTask(c *gin.Context, client *mongo.Client) {

	id, objErr := primitive.ObjectIDFromHex(c.Param("id"))
	if objErr != nil {
		c.JSON(400, util.ResMessage{
			Success: false,
			Message: "That is not a valid id",
		})
		return
	}

	taskCollection := client.Database("Todo").Collection("tasks")

	deletedDocument := models.Task{}

	deleteErr := taskCollection.FindOneAndDelete(c.Request.Context(), bson.M{"_id": id}).Decode(&deletedDocument)
	if deleteErr != nil {
		// ErrNoDocuments means that the filter did not match any documents in the collection
		if deleteErr == mongo.ErrNoDocuments {
			c.JSON(400, util.ResMessage{
				Success: false,
				Message: "There is no todo with that id",
			})
			return
		}
		c.JSON(500, util.ResError{
			Success: false,
			Error:   deleteErr,
		})
		return
	}

	c.JSON(200, util.ResMessage{
		Success: true,
		Message: "Todo deleted",
	})
}
