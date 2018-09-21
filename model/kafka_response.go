package model

import "github.com/gofrs/uuid"

// KafkaResponse is the response from consuming a Kafka message
// and operating on it. This can be used to as a "response-back"
// to indicate if the operation was successful or not.
type KafkaResponse struct {
	// AggregateID is the ID of aggregate the response is for.
	AggregateID int8 `json:"aggregate_id,omitempty"`
	// CorrelationID can be used to "identify" responses, such as checking
	// if the response if for some particular request.
	CorrelationID uuid.UUID `json:"correlation_id,omitempty"`
	// Input is the data that was being processed.
	// Use this to provide context of whatever data was attempted to be processed.
	Input []byte `json:"input"`
	// Error is the error occurred while processing the Input.
	// Convert errors to strings, this is just an indication that
	// something went wrong, so we can signal/display-error to end-
	// user. Blank Error-string means everything was fine.
	Error string `json:"error"`
	// Result is the result after an input was processed.
	// This is some data returned by processing (such as database results) etc.
	Result []byte `json:"result,omitempty"`
	// Topic is the topic on which Kafka producer should produce this message.
	// This field will not be included as part of the response, and is only
	// for referencing purposes for producer.
	Topic string `json:"-"`
}
