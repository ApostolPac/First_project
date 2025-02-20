package handlers

import (
	"GO_TUT/configs"
	"GO_TUT/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var newUser models.Human
	if err := c.BindJSON(&newUser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{ //Потому что записывает в столб, а не в строчку
			"err": err.Error(),
		})
		return
	}
	query := `INSERT INTO users (name, age, job) VALUES ($1, $2, $3) RETURNING id`
	err := configs.DB.QueryRow(query, newUser.Name, newUser.Age, newUser.Job).Scan(&newUser.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusCreated, newUser)
}

func ShowUser(c *gin.Context) {
	rows, err := configs.DB.Query("Select id, name, age, job From users")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Ошибка при получении пользователя",
		})
		return
	}
	defer rows.Close()
	var users []models.Human
	for rows.Next() {
		var user models.Human
		err := rows.Scan(&user.Id, &user.Name, &user.Age, &user.Job)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Ошибка при разборе данных",
			})
			return

		}
		users = append(users, user)
	}
	c.IndentedJSON(http.StatusOK, users)
}
