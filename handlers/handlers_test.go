package handlers

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"notetaker/models"
	"testing"
)

func TestCreateNoteHandler(t *testing.T) {
	mockHandler := NewMockNoteHandler()
	router := gin.Default()
	router.POST("/create", mockHandler.CreateNoteHandler)

	note := models.Note{ID: 1, Title: "Test Note", Text: "This is a test note"}

	payload, _ := json.Marshal(note)
	req, _ := http.NewRequest("POST", "/create", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	responseNoteJSON, err := json.Marshal(response["note"])
	assert.NoError(t, err)

	var responseNote models.Note
	err = json.Unmarshal(responseNoteJSON, &responseNote)
	assert.NoError(t, err)

	assert.Equal(t, "Note created successfully", response["message"])
	assert.Equal(t, note, responseNote)
}

func TestGetNotesHandler(t *testing.T) {
	mockHandler := NewMockNoteHandler()
	router := gin.Default()
	router.GET("/notes", mockHandler.GetNotesHandler)

	req, _ := http.NewRequest("GET", "/notes", nil)

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var notes []models.Note
	err := json.Unmarshal(w.Body.Bytes(), &notes)
	if err != nil {
		return
	}

	assert.Equal(t, 0, len(notes)) // Assuming initially no notes in mock
}
