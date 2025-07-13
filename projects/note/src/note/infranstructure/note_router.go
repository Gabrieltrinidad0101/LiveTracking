package infranstructure

import "github.com/labstack/echo"

func Router(e *echo.Echo) {
	note := e.Group("/note")
	noteController := NoteController{}

	note.GET("/list", noteController.GetAllNotes)
	note.POST("/create", noteController.CreateNote)
	note.GET("/:id", noteController.GetNote)
	note.PUT("/:id", noteController.UpdateNote)
	note.DELETE("/:id", noteController.DeleteNote)
}
