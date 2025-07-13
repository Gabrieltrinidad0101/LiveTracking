package application

type EventPublisher interface {
    Publish(eventName string, data []byte) error
}
