package internal

// TokenWrapper adds the API Token to every request.
type TokenWrapper struct {
	APIToken   string
	connection Connection
}

func NewTokenWrapper(token string) *TokenWrapper {
	return &TokenWrapper{
		APIToken: token,
	}
}

// Make adds the API Token header to every request before sending it.
func (e *TokenWrapper) Make(request *Request, passedResponse *Response) error {
	request.Header.Add("X-TrackerToken", e.APIToken)
	return e.connection.Make(request, passedResponse)
}

// Wrap wraps a connection with this wrapper.
func (e *TokenWrapper) Wrap(inner Connection) Connection {
	e.connection = inner
	return e
}
