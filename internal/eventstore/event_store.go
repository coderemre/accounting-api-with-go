package eventstore

import "context"

type EventStore interface {
    SaveEvent(ctx context.Context, event Event) error
    GetEventsByAggregateID(ctx context.Context, id string) ([]Event, error)
    GetAllEvents(ctx context.Context) ([]Event, error)
    ReplayEvents(ctx context.Context, applyFunc func(Event) error) error
}