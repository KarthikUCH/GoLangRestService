package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/KarthikUCH/buildapi/db"
	"github.com/KarthikUCH/buildapi/handler"
)

func main() {
	fmt.Println("Welcome to my first go lang REST API project")

	dbUser, dbPassword, dbName :=
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB")

	database, err := db.Initialize(dbUser, dbPassword, dbName)
	if err != nil {
		log.Fatalf("Failed connecting to database: %v", err)
	}
	defer database.Conn.Close()
	router := handler.Router(database)

	log.Fatal((http.ListenAndServe(":8080", router)))
}
