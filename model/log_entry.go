package model

// LogEntry describes the "currently-happening" event.
// Use this to show if everything is going as intended, or reasons for why not.
type LogEntry struct {
	Description   string `json:"description,omitempty"`
	ErrorCode     int    `json:"errorCode,omitempty"`
	ErrorMessage  string `json:"errorMessage,omitempty"`
	EventAction   string `json:"eventAction,omitempty"`
	ServiceAction string `json:"serviceAction,omitempty"`
	Verbosity     string `json:"verbosity,omitempty"`
}
