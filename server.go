package main

import (
	"fmt"
	"log"
	"net/http"

	"quasar/api"
	"quasar/db"
)

func main() {
	db.InitializeDatabase()
	registerHttpHandlers()
	startServer()
}

func registerHttpHandlers() {
	http.HandleFunc("/md_content/create", api.CreateMdContentHandler)
}

func startServer() {
	port := ":8080"
	fmt.Println("Server is running on port" + port)
	log.Fatal(http.ListenAndServe(port, nil))
}
