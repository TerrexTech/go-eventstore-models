package model

import (
	"time"

	"github.com/gocql/gocql"
)

// Event is an event to be stored in event-store.
type Event struct {
	// Action is the action being performed by event.
	// Examples: register_user, new_item_inventory etc.
	Action string `json:"action"`

	// AggregateID is the ID of aggregate responsible for consuming event.
	AggregateID int `json:"aggregate_id"`

	// Data is the data contained by event,
	Data string `json:"data"`

	// Timestamp is the time when the event was generated.
	Timestamp time.Time `json:"timestamp"`

	// UserID is the associated user's id who generated the event.
	UserID int `json:"user_id"`

	// UUID is the unique-indentifier for event.
	UUID gocql.UUID `json:"uuid"`

	// Version is the version for events as processed for aggregate-projection.
	// This is incremented by the aggregate itself each time an aggregate
	// projection is updated.
	Version int `json:"version"`

	// Year bucket is the year in which the event was generated.
	// This is used as the partitioning key.
	YearBucket int `json:"year_bucket"`
}
