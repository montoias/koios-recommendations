package payload

// Different types of error messages
const (
	InvalidJSON           = "invalid json"
	MissingRouteParameter = "missing route parameter in request"
	WrongParameter        = "wrong parameter"
)

// NewError returns initialized pointer to Error
func NewError(message interface{}) *Error {
	return &Error{
		Message: message,
	}
}

// Error is the response structure for all the api's if there is an error
type Error struct {
	Message interface{} `json:"error"`
}
