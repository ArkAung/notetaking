package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"notetaker/handlers"
)

func main() {
	r := gin.Default()

	noteHandler := handlers.NewNoteHandler()

	r.POST("/create", noteHandler.CreateNoteHandler)
	r.GET("/notes", noteHandler.GetNotesHandler)

	port := "8080"
	log.Printf("Server is running on port %s..\n", port)
	log.Fatal(r.Run(":" + port))
}
