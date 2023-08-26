package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"notetaker/models"
)

type NoteHandler interface {
	CreateNoteHandler(c *gin.Context)
	GetNotesHandler(c *gin.Context)
}

type noteHandler struct {
	notes map[int]models.Note
}

func NewNoteHandler() NoteHandler {
	return &noteHandler{
		notes: make(map[int]models.Note),
	}
}

func (nh *noteHandler) CreateNoteHandler(c *gin.Context) {
	var note models.Note
	if err := c.ShouldBindJSON(&note); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	note.ID = len(nh.notes) + 1
	nh.notes[note.ID] = note

	c.JSON(http.StatusCreated, gin.H{"message": "Note created", "note": note})
}

func (nh *noteHandler) GetNotesHandler(c *gin.Context) {
	notesList := make([]models.Note, 0, len(nh.notes))
	for _, note := range nh.notes {
		notesList = append(notesList, note)
	}

	c.JSON(http.StatusOK, notesList)
}
