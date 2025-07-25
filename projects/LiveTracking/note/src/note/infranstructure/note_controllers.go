package infranstructure

import (
	"fmt"
	"note/src/note/domain"

	"github.com/labstack/echo"
)

type NoteUseCase interface {
	Add(note domain.Note) error
}

type noteController struct {
	noteUseCase NoteUseCase
}

func NewNoteController(noteUseCase NoteUseCase) *noteController {
	return &noteController{
		noteUseCase,
	}
}

func (n *noteController) GetAllNotes(ctx echo.Context) error {
	return ctx.JSON(200, "Hello World")
}

func (n *noteController) CreateNote(ctx echo.Context) error {
	var note domain.Note
	ctx.Bind(&note)
	err := n.noteUseCase.Add(note)
	if err != nil {
		return err
	}
	fmt.Println(note)
	return ctx.JSON(200, "Hello World")
}

func (n *noteController) GetNote(ctx echo.Context) error {
	return ctx.JSON(200, "Hello World")
}

func (n *noteController) UpdateNote(ctx echo.Context) error {
	return ctx.JSON(200, "Hello World")
}

func (n *noteController) DeleteNote(ctx echo.Context) error {
	return ctx.JSON(200, "Hello World")
}
