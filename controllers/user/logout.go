package user

import (
	"github.com/gin-gonic/gin"
	"github.com/harekrishnamahto9872/todo-app-golang/util"
)

// Logout //
// @desc Logout the user
// @route GET /api/v1/auth/logout
// @access Private
func Logout(c *gin.Context) {

	secure := false

	c.SetCookie("token", "", 2000, "/", "", secure, true)

	c.JSON(200, util.ResMessage{
		Success: true,
		Message: "You have been logged out",
	})
}
