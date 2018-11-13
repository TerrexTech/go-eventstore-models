package model

// Command can be used to invoke a procedure in another service.
type Command struct {
	// AggregateID is the ID of aggregate the document is for.
	AggregateID int8 `json:"aggregateID,omitempty"`

	// EventAction is the action for which the command is being produced.
	EventAction string `json:"eventAction,omitempty"`

	// ServiceAction is the service-action for which the command is being produced.
	ServiceAction string `json:"serviceAction,omitempty"`

	// Data is the data required for invoking the command.
	Data []byte `json:"data,omitempty"`
}
