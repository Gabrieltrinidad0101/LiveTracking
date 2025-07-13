package infranstructure

import (
	note "note/src/note/infranstructure"

	"github.com/labstack/echo"
)

func Server() {
	e := echo.New()
	note.Router(e)
	e.Start("0.0.0.0:8000")
}
