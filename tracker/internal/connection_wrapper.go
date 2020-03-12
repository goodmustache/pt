package internal

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . ConnectionWrapper

// ConnectionWrapper can wrap a given connection allowing the wrapper to modify
// all requests going in and out of the given connection.
type ConnectionWrapper interface {
	Connection
	Wrap(innerconnection Connection) Connection
}
