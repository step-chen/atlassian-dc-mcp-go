package types

// Error represents a unified error structure
type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Details any    `json:"details,omitempty"`
}

// Error implements the error interface
func (e *Error) Error() string {
	return e.Message
}