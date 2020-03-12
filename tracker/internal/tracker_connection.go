package internal

import (
	"crypto/tls"
	"io/ioutil"
	"net"
	"net/http"
	"time"

	"github.com/goodmustache/pt/tracker/terror"
)

// ConnectionConfig is for configuring a TrackerConnection.
type ConnectionConfig struct {
	SkipSSLValidation bool
}

// TrackerConnection represents a connection to the Cloud Controller
// server.
type TrackerConnection struct {
	HTTPClient *http.Client
	UserAgent  string
}

// NewTrackerConnection returns a new TrackerConnection with provided
// configuration.
func NewTrackerConnection(config ConnectionConfig) *TrackerConnection {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: config.SkipSSLValidation,
		},
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			KeepAlive: 30 * time.Second,
			Timeout:   5 * time.Second,
		}).DialContext,
	}

	return &TrackerConnection{
		HTTPClient: &http.Client{Transport: tr},
	}
}

// Make performs the request and parses the response.
func (connection *TrackerConnection) Make(request *Request, passedResponse *Response) error {
	// In case this function is called from a retry, passedResponse may already
	// be populated with a previous response. We reset in case there's an HTTP
	// error and we don't repopulate it in populateResponse.
	passedResponse.reset()

	response, err := connection.HTTPClient.Do(request.Request)
	if err != nil {
		return err
	}

	return connection.populateResponse(response, passedResponse)
}

func (*TrackerConnection) handleStatusCodes(response *http.Response, passedResponse *Response) error {
	if response.StatusCode == http.StatusNoContent {
		passedResponse.RawResponse = []byte("{}")
	} else {
		rawBytes, err := ioutil.ReadAll(response.Body)
		defer response.Body.Close()
		if err != nil {
			return err
		}

		passedResponse.RawResponse = rawBytes
	}

	if response.StatusCode >= 400 {
		return terror.RawHTTPStatusError{
			StatusCode:  response.StatusCode,
			RawResponse: passedResponse.RawResponse,
			RequestIDs:  response.Header[http.CanonicalHeaderKey("x-request-id")],
		}
	}

	return nil
}

func (connection *TrackerConnection) populateResponse(response *http.Response, passedResponse *Response) error {
	passedResponse.HTTPResponse = response

	if resourceLocationURL := response.Header.Get("Location"); resourceLocationURL != "" {
		passedResponse.ResourceLocationURL = resourceLocationURL
	}

	err := connection.handleStatusCodes(response, passedResponse)
	if err != nil {
		return err
	}

	if passedResponse.DecodeJSONResponseInto != nil {
		err = DecodeJSON(passedResponse.RawResponse, passedResponse.DecodeJSONResponseInto)
		if err != nil {
			return err
		}
	}

	return nil
}
