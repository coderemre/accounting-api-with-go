package eventstore

import (
	"context"
	"database/sql"
)

type MySQLEventStore struct {
    db *sql.DB
}

func NewMySQLEventStore(db *sql.DB) *MySQLEventStore {
    return &MySQLEventStore{db: db}
}

func (s *MySQLEventStore) SaveEvent(ctx context.Context, e Event) error {
    _, err := s.db.ExecContext(ctx,
        `INSERT INTO events (aggregate_id, event_type, payload) VALUES (?, ?, ?)`,
        e.AggregateID, e.Type, e.Payload,
    )
    return err
}

func (s *MySQLEventStore) GetEventsByAggregateID(ctx context.Context, id string) ([]Event, error) {
    rows, err := s.db.QueryContext(ctx,
        `SELECT id, aggregate_id, event_type, payload, created_at FROM events WHERE aggregate_id = ? ORDER BY id ASC`,
        id,
    )
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var evs []Event
    for rows.Next() {
        var e Event
        if err := rows.Scan(&e.ID, &e.AggregateID, &e.Type, &e.Payload, &e.CreatedAt); err != nil {
            return nil, err
        }
        evs = append(evs, e)
    }
    return evs, rows.Err()
}

func (s *MySQLEventStore) GetAllEvents(ctx context.Context) ([]Event, error) {
    rows, err := s.db.QueryContext(ctx,
        `SELECT id, aggregate_id, event_type, payload, created_at FROM events ORDER BY id ASC`,
    )
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var evs []Event
    for rows.Next() {
        var e Event
        if err := rows.Scan(&e.ID, &e.AggregateID, &e.Type, &e.Payload, &e.CreatedAt); err != nil {
            return nil, err
        }
        evs = append(evs, e)
    }
    return evs, rows.Err()
}

func (s *MySQLEventStore) ReplayEvents(ctx context.Context, applyFunc func(Event) error) error {
    rows, err := s.db.QueryContext(ctx,
        `SELECT id, aggregate_id, event_type, payload, created_at FROM events ORDER BY id ASC`)
    if err != nil {
        return err
    }
    defer rows.Close()

    for rows.Next() {
        var e Event
        if err := rows.Scan(&e.ID, &e.AggregateID, &e.Type, &e.Payload, &e.CreatedAt); err != nil {
            return err
        }
        if err := applyFunc(e); err != nil {
            return err
        }
    }
    return rows.Err()
}