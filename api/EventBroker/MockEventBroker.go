package eventbroker

import "company-microservice/internals"

type MockEventBroker struct {
	EventsQueue []map[string]string
}

func NewMockEventBroker() *MockEventBroker {
	return &MockEventBroker{}
}

func (eb *MockEventBroker) Publish(topic string, message string) error {
	eb.EventsQueue = append(eb.EventsQueue, map[string]string{topic: message})
	internals.LogJSON(eb.EventsQueue)
	return nil
}

func (eb *MockEventBroker) Subscribe(topic string, handler func(string)) error {
	return nil
}
