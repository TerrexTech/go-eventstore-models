package model

import "github.com/TerrexTech/uuuid"

// EventStoreQuery can be used to fetch later events than the specified version.
type EventStoreQuery struct {
	// AggregateID is the id for aggregate whose events are to be fetched
	AggregateID int8 `json:"aggregate_id"`

	// AggregateVersion is the highest version of events that have been
	// already fetched by the aggregate. The event-store will be queried
	// for events greater than this version.
	AggregateVersion int64 `json:"aggregate_version"`

	// CorrelationID can be used to "identify" responses, such as checking
	// if the response is for some particular request.
	// Including CorrelationID will result in inclusion of this ID in any
	// responses generated as per result of event's processing.
	CorrelationID uuuid.UUID `json:"correlation_id,omitempty"`
}
