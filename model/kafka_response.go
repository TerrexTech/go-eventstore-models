package model

// KafkaResponse is the response from consuming a Kafka message
// and operating on it. This can be used to as a "response-back"
// to indicate if the operation was successful or not.
type KafkaResponse struct {
	// AggregateID is the ID of aggregate the response is for.
	AggregateID int8 `json:"aggregate_id,omitempty"`
	// CorrelationID can be used to "identify" responses, such as checking
	// if the response if for some particular request.
	CorrelationID int8 `json:"correlation_id,omitempty"`
	// Input is the message-input received by Consumer.
	// Use this to provide context of whatever data was attempted to be processed.
	Input string `json:"input"`
	// Error is the error occurred while processing the Input.
	// Convert errors to strings, this is just an indication that
	// something went wrong, so we can signal/display-error to end-
	// user. Blank Error-string means everything was fine.
	Error string `json:"error"`
	// Result is the result after an input was processed.
	// This is some data returned by processing (such as database results) etc.
	Result string `json:"result,omitempty"`
}
