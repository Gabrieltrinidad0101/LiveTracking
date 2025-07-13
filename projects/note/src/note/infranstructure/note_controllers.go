package infranstructure

import (
	"github.com/labstack/echo"
)

type NoteController struct{}

func (n *NoteController) GetAllNotes(ctx echo.Context) error {
	return ctx.JSON(200, "Hello World")
}

func (n *NoteController) CreateNote(ctx echo.Context) error {
	return ctx.JSON(200, "Hello World")
}

func (n *NoteController) GetNote(ctx echo.Context) error {
	return ctx.JSON(200, "Hello World")
}

func (n *NoteController) UpdateNote(ctx echo.Context) error {
	return ctx.JSON(200, "Hello World")
}

func (n *NoteController) DeleteNote(ctx echo.Context) error {
	return ctx.JSON(200, "Hello World")
}
