package model

// KafkaResponse is the response from consuming a Kafka message
// and operating on it. This can be used to as a "response-back"
// to indicate if the operation was successful or not.
type KafkaResponse struct {
	// AggregateID is the ID of aggregate the response is for.
	// This can be "-1" if the aggregateID at that point is unknown.
	AggregateID int8 `json:"aggregate_id"`
	// Input is the message-input received by Consumer.
	Input string `json:"input"`
	// Error is the error occurred while processing the Input.
	// Convert errors to strings, this is just an indication that
	// something went wrong, so we can signal/display-error to end-
	// user. Blank Error-string means everything was fine.
	Error string `json:"error"`
}
