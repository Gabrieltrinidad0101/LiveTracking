package infranstructure

import (
	"log"
	"note/src/note/application"
	"note/src/note/infranstructure/rabbitmq"
	"note/src/note/infranstructure/tools"

	"github.com/labstack/echo"
	"github.com/streadway/amqp"
)

func Router(e *echo.Echo, conn *amqp.Connection) {
	note := e.Group("/note")

	jsonToBytes := tools.JsonToBytes
	noteRabbit, err := rabbitmq.NewRabbitMQ(conn, "note", "topic", jsonToBytes)
	if err != nil {
		log.Fatalf("Error al crear el publisher: %v", err)
	}
	noteUseCase := application.NewNoteUseCase(noteRabbit)
	noteController := NewNoteController(noteUseCase)

	note.GET("/list", noteController.GetAllNotes)
	note.POST("/create", noteController.CreateNote)
	note.GET("/:id", noteController.GetNote)
	note.PUT("/:id", noteController.UpdateNote)
	note.DELETE("/:id", noteController.DeleteNote)
}
