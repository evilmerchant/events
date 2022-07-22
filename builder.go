package events

import (
	"log"
	"time"

	"github.com/google/uuid"
)

type EventBuilder struct {
	event *Event
}

func New(eventType string) *EventBuilder {
	builder := &EventBuilder{
		event: &Event{
			Id:        uuid.New(),
			Type:      eventType,
			Timestamp: time.Now().UTC(),
		},
	}
	return builder
}

func (builder *EventBuilder) WithPayload(payload any) *EventBuilder {
	builder.event.Payload = payload
	return builder
}

func (builder *EventBuilder) WithStreamId(streamId string) *EventBuilder {
	builder.event.Stream = Stream{
		Id: streamId,
	}
	return builder
}

func (builder *EventBuilder) WithType(streamId string) *EventBuilder {
	builder.event.Stream = Stream{
		Id: streamId,
	}
	return builder
}

func (builder *EventBuilder) Event() *Event {

	if builder.event.Type == "" {
		log.Fatalln("event type is not set")
	}
	if builder.event.Stream.Id == "" {
		log.Fatalln("stream id is not set")
	}

	return builder.event
}
