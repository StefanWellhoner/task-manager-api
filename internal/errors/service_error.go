package errors

type Type string

const (
	ValidationError   Type = "VALIDATION_ERROR"
	UnauthorizedError Type = "UNAUTHORIZED_ERROR"
	NotFoundError     Type = "NOT_FOUND_ERROR"
	InternalError     Type = "INTERNAL_ERROR"
	ConflictError     Type = "CONFLICT_ERROR"
)

type ServiceError struct {
	Type    Type   `json:"type"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func (e *ServiceError) Error() string {
	return e.Message
}

func NewServiceError(errorType Type, message string, status int) *ServiceError {
	return &ServiceError{
		Type:    errorType,
		Message: message,
		Status:  status,
	}
}
