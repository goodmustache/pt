package internal

//counterfeiter:generate . Connection

// Connection creates and executes http requests
type Connection interface {
	Make(request *Request, passedResponse *Response) error
}
