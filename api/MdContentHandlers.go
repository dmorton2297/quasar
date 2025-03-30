package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"quasar/db"
)

type CreateMdContent struct {
	Content string `json:"content"`
	Name    string `json:"name"`
}

func CreateMdContentHandler(w http.ResponseWriter, r *http.Request) {
	fakeUserId := "00000000-0000-0000-0000-000000000000"
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method allowed", http.StatusMethodNotAllowed)
		return
	}

	var input CreateMdContent
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Create new MdContent instance
	newMdContent := db.MdContent{
		Name:            input.Name,
		Content:         input.Content,
		Hidden:          false, // Default value
		CreatedByUserId: fakeUserId,
	}

	// Insert into the database
	result := db.Connection.Create(&newMdContent)
	if result.Error != nil {
		http.Error(w, "Failed to create record", http.StatusInternalServerError)
		return
	}

	// Create mdContent message
	createMessage := db.MdContentCreated{
		CreatedEntity:   newMdContent,
		CreatedByUserId: fakeUserId,
	}

	// Encode message to json
	messageJsonBytes, err := json.Marshal(createMessage)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// For now just print the message
	log.Printf("Message: %s\n", string(messageJsonBytes))

	// Return the created record
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newMdContent)
}
