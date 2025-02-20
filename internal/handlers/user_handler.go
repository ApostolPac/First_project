package handlers

import (
	"GO_TUT/internal/models"
	"GO_TUT/internal/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var newUser models.Human
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}
	query := `INSERT INTO users (name, age, job) VALUES ($1, $2, $3) RETURNING id`
	err := repository.DB.QueryRow(query, &newUser.Name, &newUser.Age, &newUser.Job).Scan(&newUser.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusCreated, newUser)
}

func ShowUser(c *gin.Context) {
	rows, err := repository.DB.Query("Select id, name, age, job From users")
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
	c.JSON(http.StatusOK, users)
}
func ChangeOneUser(c *gin.Context) {
	var id string
	id = c.Param("id")
	var newUser models.Human
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": err.Error(),
		})
		return
	}
	_, err := repository.DB.Exec("UPDATE users Set name = $1, age = $2, job = $3 Where id = $4", &newUser.Name, &newUser.Age, &newUser.Job, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, "Changed")
}
func DeleteOneUser(c *gin.Context) {
	var id string
	id = c.Param("id")
	_, err := repository.DB.Exec("Delete From users Where id = $1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, "Deleted")
}

func ShowOneUser(c *gin.Context) {
	var id string
	id = c.Param("id")

	rows := repository.DB.QueryRow("Select id, name, age, job From users Where id = $1", id)
	var users models.Human
	err := rows.Scan(&users.Id, &users.Name, &users.Age, &users.Job)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusOK, users)
}
