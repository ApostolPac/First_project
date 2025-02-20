package handlers

import (
	"net/http"

	"GO_TUT/internal/models"

	"github.com/gin-gonic/gin"
)

var users = []models.Human{}

func CreateUser(c *gin.Context) {
	var newUser models.Human
	if err := c.BindJSON(&newUser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}
	users = append(users, newUser)

	c.JSON(http.StatusOK, users)
}

func ShowUser(c *gin.Context) {
	if len(users) == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"error": "No user created yet",
		})
		return
	}
	c.IndentedJSON(http.StatusOK, users)
}
