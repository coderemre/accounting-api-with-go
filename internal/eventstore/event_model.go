package eventstore

import (
	"encoding/json"
	"time"
)

type Event struct {
    ID          int64           `json:"id"`
    AggregateID string          `json:"aggregate_id"`
    Type        string          `json:"type"`
    Payload     json.RawMessage `json:"payload"`
    CreatedAt   time.Time       `json:"created_at"`
}