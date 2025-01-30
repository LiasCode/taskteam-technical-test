package eventbroker

type EventBroker interface {
	Publish(topic string, message string) error
	Subscribe(topic string, handler func(string)) error
}
