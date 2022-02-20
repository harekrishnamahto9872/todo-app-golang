package main

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/harekrishnamahto9872/todo-app-golang/config"
)

func main() {

	client := config.ConnectDB()

	router := gin.Default()

	router.Run()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	defer client.Disconnect(ctx)

}
