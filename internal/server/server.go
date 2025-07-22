package server

import (
	"fmt"
	"log"
	"net/http"

	"notesapp/internal/models"
	"notesapp/internal/handlers"

)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from your first Go server!")
}

func Start() {
	err := models.InitDB()
	if err != nil {
		log.Fatalf("Failed to connect to DB : %v", err)
	}

	http.HandleFunc("/", handler)
	http.HandleFunc("/notes", handlers.NotesHandler)
	http.HandleFunc("/notes/", handlers.DeleteNoteHandler)


	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
