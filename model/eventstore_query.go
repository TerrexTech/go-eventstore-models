package model

// EventStoreQuery can be used to fetch later events than the specified version.
type EventStoreQuery struct {
	// AggregateID is the id for aggregate whose events are to be fetched
	AggregateID int8 `json:"aggregate_id"`
	// AggregateVersion is the highest version of events that have been
	// already fetched by the aggregate. The event-store will be queried
	// for events greater than this version.
	AggregateVersion int64 `json:"aggregate_version"`
	// YearBucket is the partitioning-key for the Cassandra table.
	// Specify this to let the query-handler know which partition to
	// user for query operations.
	YearBucket int16 `json:"year_bucket"`
}
