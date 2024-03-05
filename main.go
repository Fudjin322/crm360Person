package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"sqlBoiler/models"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:root@127.0.0.1:5432/postgres?sslmode=disable"
)

func InitDB() (*sqlx.DB, error) {
	db, err := sqlx.Connect(dbDriver, dbSource)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func main() {
	var err error
	db, err := InitDB()
	if err != nil {
		panic(err)
	}

	db, err = sqlx.Connect(dbDriver, dbSource)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	r.GET("/persons", models.GetAllPersons)
	r.POST("/persons", models.CreatePerson)

	port := 8080
	address := fmt.Sprintf(":%d", port)
	log.Printf("Сервер запущен на порту %d", port)
	r.Run(address)
}
