package responses

type FailedResponseBody struct {
	Message string `json:"message"`
}

// A ValidationError is an error that is used when the required input fails validation.
// swagger:response ValidationError
type ValidationError struct {
	// The error message
	// in: body
	Payload Body
}
