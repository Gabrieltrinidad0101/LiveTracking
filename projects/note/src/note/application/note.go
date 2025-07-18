package application

import (
	"note/src/note/domain"
)

type EventBus interface {
	Publish(routingKey string, body interface{}) error
}

type NoteUseCase struct {
	eventBus EventBus
}

func NewNoteUseCase(eventBus EventBus) *NoteUseCase {
	return &NoteUseCase{
		eventBus,
	}
}

func (n *NoteUseCase) Add(note domain.Note) error {
	err := n.eventBus.Publish("note:create", note)
	return err
}
