package container

import (
	event_broker_implementation "company-microservice/api/EventBroker"
	eventbroker "company-microservice/pkg/EventBroker"
)

func (c *Container) GetEventBroker() eventbroker.EventBroker {
	if c.eventBroker != nil {
		return c.eventBroker
	}

	mockEventBroker := event_broker_implementation.NewMockEventBroker()

	if mockEventBroker == nil {
		panic("MockEventBroker is nil")
	}

	c.eventBroker = mockEventBroker

	return c.eventBroker
}
