package model

import (
	"time"

	"github.com/TerrexTech/uuuid"
)

// Event refers to a change in system, and is stored in event-store.
type Event struct {
	// Action is the action being performed by event.
	// Examples: register_user, new_item_inventory etc.
	Action string `cql:"action" json:"action"`

	// CorrelationID can be used to "identify" responses, such as checking
	// if the response is for some particular request.
	// Including CorrelationID will result in inclusion of this ID in any
	// responses generated as per result of event's processing.
	CorrelationID uuuid.UUID `json:"correlation_id,omitempty"`

	// AggregateID is the ID of aggregate responsible for consuming event.
	AggregateID int8 `cql:"aggregate_id" json:"aggregate_id"`

	// Data is the data contained by event.
	Data []byte `cql:"data" json:"data"`

	// Timestamp is the time when the event was generated.
	Timestamp time.Time `cql:"timestamp" json:"timestamp"`

	// UserUUID is the V4-UUID of the user who generated the event.
	UserUUID uuuid.UUID `cql:"user_uuid" json:"user_uuid"`

	// TimeUUID is the V1-UUID unique-indentifier for event.
	TimeUUID uuuid.UUID `cql:"time_uuid" json:"time_uuid"`

	// Version is the version for events as processed for aggregate-projection.
	// This is incremented by the aggregate itself each time it updates its
	// projection.
	Version int64 `cql:"version" json:"version"`

	// Year bucket is the year in which the event was generated.
	// This is used as the partitioning key.
	YearBucket int16 `cql:"year_bucket" json:"year_bucket"`
}
