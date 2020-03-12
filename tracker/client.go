package tracker

import (
	"fmt"
	"runtime"

	"github.com/goodmustache/pt/tracker/internal"
	"github.com/tedsuo/rata"
)

const TRACKER_URL = "https://www.pivotaltracker.com/services/v5"

// Client can be used to communicate with the Tracker API.
type Client struct {
	connection internal.Connection
	router     *rata.RequestGenerator
	userAgent  string

	clock internal.Clock
}

// Config allows the Client to be configured
type Config struct {
	// AppName is the name of the application/process using the client.
	AppName string

	// AppVersion is the version of the application/process using the client.
	AppVersion string

	// APIToken is the tracker API token to make requests with.
	APIToken string

	// Wrappers that apply to the client connection.
	Wrappers []internal.ConnectionWrapper

	// SkipSSLValidation will skip hostname validation when set to true.
	SkipSSLValidation bool
}

// NewClient returns a new Client.
func NewClient(config Config) *Client {
	return createClient(TRACKER_URL, new(internal.RealTime), config)
}

// TestClient returns a new client explicitly meant for internal testing. This
// should not be used for production code.
func TestClient(rootURL string, clock internal.Clock, config Config) *Client {
	return createClient(rootURL, clock, config)
}

// WrapConnection wraps the current Client connection in the wrapper.
func (client *Client) WrapConnection(wrapper internal.ConnectionWrapper) {
	client.connection = wrapper.Wrap(client.connection)
}

func createClient(rootURL string, clock internal.Clock, config Config) *Client {
	userAgent := fmt.Sprintf("%s/%s (%s; %s %s)", config.AppName, config.AppVersion, runtime.Version(), runtime.GOARCH, runtime.GOOS)
	client := &Client{
		clock:      clock,
		connection: internal.NewTrackerConnection(internal.ConnectionConfig{SkipSSLValidation: config.SkipSSLValidation}),
		router:     rata.NewRequestGenerator(rootURL, internal.APIRoutes),
		userAgent:  userAgent,
	}

	wrappers := append([]internal.ConnectionWrapper{
		internal.NewErrorWrapper(),
		internal.NewTokenWrapper(config.APIToken),
	}, config.Wrappers...)
	for _, wrapper := range wrappers {
		client.connection = wrapper.Wrap(client.connection)
	}

	return client
}
