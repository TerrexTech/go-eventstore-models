package model

// EventMeta contains the information about hydrating Aggregate-Projections.
// The AggregateVersion tracks the version to be used by new events for that
// aggregate. This difference between the aggregate-projection version and
// AggregateVersion gives us the versions of the events yet to be applied to the
// aggregate projection.
type EventMeta struct {
	// AggregateVersion tracks the version to be used
	// by new events for that aggregate.
	AggregateVersion int64 `cql:"aggregate_version" json:"aggregate_version"`
	// AggregateID corresponds to AggregateID in
	// event-store and ID in aggregate-projection.
	AggregateID int8 `cql:"aggregate_id" json:"aggregate_id"`
	// PartitionKey is the partitioning key for events_meta table.
	PartitionKey int8 `cql:"partition_key" json:"partition_key"`
}
