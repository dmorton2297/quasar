package main

import (
	"fmt"
	"log"
	"net/http"

	"quasar/db"
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Server is running\n"))
}

func main() {
	db.InitializeDatabase()
	http.HandleFunc("/status", pingHandler)

	port := ":8080"
	fmt.Println("Server is running on port" + port)
	log.Fatal(http.ListenAndServe(port, nil))
}
