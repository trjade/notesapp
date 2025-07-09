package handlers

import (
	"bytes"
"net/http"
"net/http/httptest"
"testing"

)

func TestCreateNoteValidation(t *testing.T){
	req := httptest.NewRequest(http.MethodPost, "/notes", bytes.NewBuffer([]byte(`{"title":"", "content":""}`)))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	NotesHandler(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", w.Code)
	}
}