package main

import (
	"GO_TUT/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	r.POST("/ping", handlers.CreateUser)

	r.GET("/ping", handlers.ShowUser)

	r.Run(":8080")

}
