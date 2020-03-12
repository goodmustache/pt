package internal

// ErrorWrapper is the wrapper that converts responses with 4xx and 5xx status
// codes to an error.
type ErrorWrapper struct {
	connection Connection
}

func NewErrorWrapper() *ErrorWrapper {
	return new(ErrorWrapper)
}

// Make creates a connection in the wrapped connection and handles errors
// that it returns.
func (e *ErrorWrapper) Make(request *Request, passedResponse *Response) error {
	err := e.connection.Make(request, passedResponse)
	return err
}

// Wrap wraps a connection in this error handling wrapper.
func (e *ErrorWrapper) Wrap(inner Connection) Connection {
	e.connection = inner
	return e
}
