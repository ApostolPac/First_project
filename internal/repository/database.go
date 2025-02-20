package repository

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDb() {
	var err error
	conn := "host=localhost port=5432 user=postgres password=1234 dbname=postgres sslmode=disable"
	DB, err = sql.Open("postgres", conn)
	if err != nil {
		log.Fatal("Ошибка подключения к дб", err)
	}
	err = DB.Ping() // проверяя подключение
	if err != nil {
		log.Fatal("Бд не доступна", err)
	}
	fmt.Println("Успешное подключение")
}
