package main

import (
	"GO_TUT/configs"
	"GO_TUT/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {

	configs.ConnectDb() //Подключение

	r := gin.Default()

	r.POST("/users", handlers.CreateUser) //Ввод данных о пользователе

	r.GET("/users", handlers.ShowUser) //Вывод данных о пользователе

	r.Run(":8080")

}
