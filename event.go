package events

import (
	"time"

	"github.com/google/uuid"
)

type Stream struct {
	Id string `json:"id"`
}

type Event struct {
	Self      string    `json:"_self,omitempty"`
	Id        uuid.UUID `json:"id"`
	Stream    Stream    `json:"stream"`
	Type      string    `json:"type"`
	Payload   any       `json:"payload"`
	Timestamp time.Time `json:"timestamp"`
}
