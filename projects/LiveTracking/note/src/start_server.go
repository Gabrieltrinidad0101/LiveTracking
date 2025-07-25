package src

import (
	note "note/src/note/infranstructure"

	"github.com/labstack/echo"
	"github.com/streadway/amqp"
)

func Server(conn *amqp.Connection) {
	e := echo.New()
	note.Router(e, conn)
	e.Start("0.0.0.0:8000")
}
