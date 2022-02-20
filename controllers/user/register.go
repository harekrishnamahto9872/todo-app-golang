package user

import (
	"fmt"
	"log"

	"github.com/harekrishnamahto9872/todo-app-golang/util"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/harekrishnamahto9872/todo-app-golang/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Register //
// @desc Create a new user from User struct and add it to the database
// @route POST /api/v1/auth/register
// @access Public
func Register(c *gin.Context, client *mongo.Client) {

	user := models.UserCred{}

	bindErr := c.ShouldBindJSON(&user)
	fmt.Println("Value of c", c)
	if bindErr != nil {
		fmt.Println("some binding error")
		log.Fatal(bindErr)
	}

	// checks to make sure that all fields are populated
	if user.Name == "" {
		c.JSON(400, util.ResMessage{
			Success: false,
			Message: "Please add a name",
		})
		return
	} else if user.Email == "" {
		c.JSON(400, util.ResMessage{
			Success: false,
			Message: "Please add an email",
		})
		return
	} else if user.Password == "" {
		c.JSON(400, util.ResMessage{
			Success: false,
			Message: "Please add a password",
		})
		return
	}

	user.Encrypt(user.Password)

	usersCollection := client.Database("Todo").Collection("users")

	alreadyExists := models.User{}

	// search db to make sure provided email is unique
	findOneErr := usersCollection.FindOne(c.Request.Context(), bson.M{"email": user.Email}).Decode(&alreadyExists)
	if findOneErr == nil {
		c.JSON(400, util.ResMessage{
			Success: false,
			Message: "That email already exists",
		})
		return
	}
	fmt.Println("context:", c.Request.Context())
	insertOneResult, insertErr := usersCollection.InsertOne(c.Request.Context(), user)
	if insertErr != nil {
		fmt.Println("inserterr binding error", insertErr)
		c.JSON(400, util.ResError{
			Success: false,
			Error:   insertErr,
		})
	} else {
		token, getSignedErr := user.GetSignedJWT(insertOneResult.InsertedID.(primitive.ObjectID).Hex())
		if getSignedErr != nil {
			c.JSON(400, util.ResError{
				Success: false,
				Error:   getSignedErr,
			})
			return
		}

		// secure cookie unless in development env
		secure := false

		c.SetCookie("token", token, 2000, "/", "", secure, true)

		c.JSON(200, util.ResUser{
			Success: true,
			Message: models.UserRes{
				Name:  user.Name,
				Email: user.Email,
			},
		})
	}
}
