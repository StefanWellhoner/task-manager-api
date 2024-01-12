package model

// Error Model
//
// Error model is used to return error messages to the client.
//
// swagger:model Error
type Error struct {
	// The general error message
	//
	// required: true
	// example: Unauthorized
	Error string `json:"error"`
	// The http error code
	//
	// required: true
	// example: 401
	ErrorCode int `json:"errorCode"`
	// The error description
	//
	// required: true
	// example: You are not authorized to access this resource
	ErrorDescription string `json:"errorDescription"`
}
