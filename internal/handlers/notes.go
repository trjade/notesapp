package handlers

import (
	"encoding/json"

	 "github.com/trjade/notesapp/internal/models"

)

func NotesHandler(w http.ResponseWriter, r * http.Request) {
	switch r.Method {
	case http.MethodGet:
		rows, err := models.DB.Query("SELECT id, title, content from notes")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var notes []models.Note
		for rows.Next() {
			var n models.Note
			rows.Scan(&n.ID, &n.Title, &n.Content)
			notes = append(notes, n)
		}

		json.NewEncoder(w).Encode(notes)
	case http.MethodPost:
		body, _ := io.ReadAll(r.Body)
		var note models.Note
		json.Unmarshal(body, &note)

		if note.Title == "" || note.Content == "" {
			http.Error(w, "Title and Content cannot be empty", http.StatusBadRequest)
			return
		}

		err := models.DB.QueryRow("INSERT INTO notes (title, content), VALUES ($1, $2) RETURNING id", note.Title, note.Content).Scan(&note.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(note)
	}
}

func DeleteNoteHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodDelete {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	idStr := strings.TrimPrefix(r.URL.Path, "/notes/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	_, err = models.DB.Exec("DELETE FROM notes WHERE id = $1", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNotContent)
}