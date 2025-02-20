package main

import (
	"GO_TUT/internal/handlers"
	"GO_TUT/internal/repository"

	"github.com/gin-gonic/gin"
)

func main() {

	repository.ConnectDb() //Подключение

	r := gin.Default()

	r.POST("/users", handlers.CreateUser) //Ввод данных о пользователе

	r.GET("/users", handlers.ShowUser) //Вывод данных о пользователе
	r.GET("/users/:id", handlers.ShowOneUser)
	r.PUT("/users/:id", handlers.ChangeOneUser)
	r.DELETE("/users/:id",handlers.DeleteOneUser)
	r.Run(":8080")

}
