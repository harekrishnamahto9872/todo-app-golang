package task

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/harekrishnamahto9872/todo-app-golang/models"
	"github.com/harekrishnamahto9872/todo-app-golang/util"
	"github.com/mitchellh/mapstructure"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// UpdateTask //
// @desc Updates a task with id in request params and data in body
// @route PUT /api/v1/tasks/:id
// @access Private
func UpdateTask(c *gin.Context, client *mongo.Client) {

	updateFields := bson.M{
		"updatedAt": time.Now().Unix(),
	}

	bindErr := c.ShouldBindJSON(&updateFields)
	if bindErr != nil {
		c.JSON(400, util.ResError{
			Success: false,
			Error:   bindErr,
		})
		return
	}

	// sets the operator for the fields to be updated
	update := bson.M{"$set": updateFields}

	// new ObjectID from request params.  this is the id of the document to be updated
	id, objErr := primitive.ObjectIDFromHex(c.Param("id"))
	if objErr != nil {
		c.JSON(400, util.ResMessage{
			Success: false,
			Message: "That is not an id",
		})
	}

	taskCollection := client.Database("Todo").Collection("tasks")

	// this will be the document BEFORE it has been updated
	// this must be a map so that I can index it and update the values because I want to return the updated document
	updated := bson.M{}

	// find one by id and update
	updateErr := taskCollection.FindOneAndUpdate(c.Request.Context(), bson.M{"_id": id}, update).Decode(&updated)
	if updateErr != nil {
		// ErrNoDocuments means that the filter did not match any documents in the collection
		if updateErr == mongo.ErrNoDocuments {
			c.JSON(400, util.ResMessage{
				Success: false,
				Message: "There is no todo with that id",
			})
			return
		}
		c.JSON(500, util.ResError{
			Success: false,
			Error:   updateErr,
		})
	}

	// merge the updated fields with the document that has been returned form the update operation
	for key, value := range updateFields {
		updated[key] = value
	}

	// the response struct takes models.Task{} so this function will fill it will the updated document
	updatedTodo := models.Task{}
	updatedTodo.ID = updated["_id"].(primitive.ObjectID)
	decodeErr := mapstructure.Decode(updated, &updatedTodo)
	if decodeErr != nil {
		c.JSON(500, util.ResError{
			Success: false,
			Error:   decodeErr,
		})
	}

	c.JSON(200, util.ResTask{
		Success: true,
		Message: updatedTodo,
	})
}
