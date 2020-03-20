package internal

//counterfeiter:generate . ConnectionWrapper

// ConnectionWrapper can wrap a given connection allowing the wrapper to modify
// all requests going in and out of the given connection.
type ConnectionWrapper interface {
	Connection
	Wrap(innerconnection Connection) Connection
}
