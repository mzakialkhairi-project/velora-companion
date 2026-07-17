// Package event provides event types for domain events.
package event

import "time"

// Event is the base interface for all domain events
type Event interface {
	// EventType returns the type of the event
	EventType() string

	// OccurredAt returns the time when the event occurred
	OccurredAt() time.Time
}

// BaseEvent provides common fields for all domain events
type BaseEvent struct {
	occurredAt time.Time
}

// NewBaseEvent creates a new BaseEvent with the current time
func NewBaseEvent() BaseEvent {
	return BaseEvent{
		occurredAt: time.Now(),
	}
}

// OccurredAt implements Event
func (e BaseEvent) OccurredAt() time.Time {
	return e.occurredAt
}

// DomainEvent is the marker interface for domain events
type DomainEvent interface {
	Event
}

// IntegrationEvent is the marker interface for integration events
type IntegrationEvent interface {
	Event
}
