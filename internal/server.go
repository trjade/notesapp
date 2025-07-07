package server

import (
	"fmt"
	"log"
	"net/http"

)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from your first Go server!")
}

func Start() {
	http.HandleFunc("/", handler)
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
