package application

type EventBus interface {
	Publish(routingKey string, body []byte) error
}

type NoteUseCase struct {
	eventBus EventBus
}

func NewNoteUseCase(eventBus EventBus) *NoteUseCase {
	return &NoteUseCase{
		eventBus,
	}
}

func (n *NoteUseCase) Add() {
	n.eventBus.Publish("note", []byte{})
}
