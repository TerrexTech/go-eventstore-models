package model

import "github.com/TerrexTech/uuuid"

// EventStoreQuery can be used to fetch later events than the specified version.
type EventStoreQuery struct {
	// AggregateID is the id for aggregate whose events are to be fetched
	AggregateID int8 `json:"aggregateID,omitempty"`

	// AggregateVersion is the highest version of events that have been
	// already fetched by the aggregate. The event-store will be queried
	// for events greater than this version.
	AggregateVersion int64 `json:"aggregateVersion,omitempty"`

	// CorrelationID can be used to "identify" responses, such as checking
	// if the response is for some particular request.
	// Including CorrelationID will result in inclusion of this ID in any
	// responses generated as per result of event's processing.
	CorrelationID uuuid.UUID `json:"correlationID,omitempty"`

	// EventAction is the core-action being performed by event.
	// For example, "insert" is EventAction, but "insertUser" is ServiceAction,
	// informing service that a user was inserted.
	EventAction string `json:"eventAction,omitempty"`

	// ServiceAction is the service-specific Action for the event.
	// For example, "insert" is EventAction, but "insertUser" is ServiceAction,
	// informing service that a user was inserted.
	ServiceAction string `json:"serviceAction,omitempty"`

	// YearBucket is the partition-key for Event-Table.
	YearBucket int16 `json:"yearBucket,omitempty"`

	// UUID is the V4-UUID Query-Identifier.
	// This can be used to "identify" responses.
	UUID uuuid.UUID `json:"uuid,omitempty"`
}
