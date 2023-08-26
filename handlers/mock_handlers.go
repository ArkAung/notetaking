package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"notetaker/models"
)

type MockNoteHandler struct {
	notes map[int]models.Note
}

func NewMockNoteHandler() NoteHandler {
	return &MockNoteHandler{
		notes: make(map[int]models.Note),
	}
}

func (mnh *MockNoteHandler) CreateNoteHandler(c *gin.Context) {
	var note models.Note
	if err := c.ShouldBindJSON(&note); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Simulate adding the note to the mock data
	note.ID = len(mnh.notes) + 1
	mnh.notes[note.ID] = note

	c.JSON(http.StatusCreated, gin.H{"message": "Note created successfully", "note": note})
}

func (mnh *MockNoteHandler) GetNotesHandler(c *gin.Context) {
	// Simulate fetching notes from the mock data
	c.JSON(http.StatusOK, mnh.notes)
}
