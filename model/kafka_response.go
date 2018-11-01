package model

import "github.com/TerrexTech/uuuid"

// KafkaResponse is the response from consuming a Kafka message
// and operating on it. This can be used to as a "response-back"
// to indicate if the operation was successful or not.
type KafkaResponse struct {
	// AggregateID is the ID of aggregate the response is for.
	AggregateID int8 `json:"aggregate_id,omitempty"`

	// Action is the action corresponding to which the response was produced.
	Action string `json:"action,omitempty"`

	// CorrelationID can be used to "identify" responses, such as checking
	// if the response is for some particular request.
	// Including CorrelationID will result in inclusion of this ID in any
	// responses generated as per result of event's processing.
	CorrelationID uuuid.UUID `json:"correlation_id,omitempty"`

	// Error is the error occurred while processing the Input.
	// Convert errors to strings, this is just an indication that
	// something went wrong, so we can signal/display-error to end-
	// user. Blank Error-string means everything was fine.
	Error string `json:"error"`

	// ErrorCode can be used to identify type of error.
	ErrorCode int16 `json:"error_code,omitempty"`

	// Input is the data that was being processed.
	// Use this to provide context of whatever data was attempted to be processed.
	Input []byte `json:"input,omitempty"`

	// Result is the result after an input was processed.
	// This is some data returned by processing (such as database results) etc.
	Result []byte `json:"result,omitempty"`

	// Topic is the topic on which Kafka producer should produce this message.
	// This field will not be included as part of the response, and is only
	// for referencing purposes for producer.
	Topic string `json:"-"`

	// UUID is the V4-UUID Response-Identifier.
	UUID uuuid.UUID `json:"uuid,omitempty"`
}
