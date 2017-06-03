package payload

// NewResponse returns initialized pointer to Response
func NewResponse(message interface{}) *Response {
	return &Response{
		Message: message,
	}
}

// Response is the general response structure for the APIs
type Response struct {
	Message interface{} `json:"message"`
}
