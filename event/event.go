package event

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

var ErrEventUnmarshalFailed = errors.New("unmarshal of event failed")

type EventType string

// Event is a domain event marker.
type Event interface {
	EventType() EventType

	// The event ID is only used to maintain a strict order
	// in the event store, and as such only set when the
	// event is stored. For new events the command will not
	// have access to the id and should not use it.
	// For ordering the event date should be used instead.
	SetID(uint64)
	GetID() uint64

	SetDate(time.Time)
	GetDate() time.Time
}

func (et EventType) Unmarshal(data []byte) (Event, error) {
	var target Event
	switch et {
	case PlayerCreatedType:
		target = &PlayerCreated{}
	case GameCreatedType:
		target = &GameCreated{}
	case ProjectCreatedType:
		target = &ProjectCreated{}
	case MatchCreatedType:
		target = &MatchCreated{}
	}

	if err := json.Unmarshal(data, &target); err != nil {
		return nil, fmt.Errorf("unmarshal failed: %w", err)
	}

	return target, nil
}

// EventMetaData contains common meta data for Events.
type EventMetaData struct {
	ID   uint64    `json:"id"`
	Date time.Time `json:"date"`
}

func (md *EventMetaData) SetID(id uint64) {
	md.ID = id
}

func (md *EventMetaData) GetID() uint64 {
	return md.ID
}

func (md *EventMetaData) SetDate(date time.Time) {
	md.Date = date
}

func (md *EventMetaData) GetDate() time.Time {
	return md.Date
}
