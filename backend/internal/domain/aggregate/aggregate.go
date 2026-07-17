// Package aggregate provides aggregate types for the domain layer.
package aggregate

import (
	"sync"

	"github.com/mzakiaklhairi/velora/internal/domain/entity"
	"github.com/mzakiaklhairi/velora/internal/domain/event"
)

// AggregateRoot is the base type for all aggregate roots
type AggregateRoot struct {
	entity.BaseEntity
	mu      sync.RWMutex
	events  []event.DomainEvent
	version int
}

// NewAggregateRoot creates a new AggregateRoot
func NewAggregateRoot() *AggregateRoot {
	return &AggregateRoot{
		events: make([]event.DomainEvent, 0),
	}
}

// AddEvent adds a domain event to the aggregate
func (a *AggregateRoot) AddEvent(evt event.DomainEvent) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.events = append(a.events, evt)
}

// PullEvents returns and clears all domain events
func (a *AggregateRoot) PullEvents() []event.DomainEvent {
	a.mu.Lock()
	defer a.mu.Unlock()

	events := a.events
	a.events = make([]event.DomainEvent, 0)
	return events
}

// Version returns the aggregate version
func (a *AggregateRoot) Version() int {
	return a.version
}

// IncrementVersion increments the aggregate version
func (a *AggregateRoot) IncrementVersion() {
	a.version++
}
