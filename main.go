package main

import (
	"fmt"
	"go-task-tracker/config"
	"go-task-tracker/routes"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()

	router := routes.SetupRoutes()

	fmt.Println("🚀 Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
