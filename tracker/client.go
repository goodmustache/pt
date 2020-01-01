// Package tracker is designed to work with the Pivotal Tracker V5 API:
// https://www.pivotaltracker.com/help/api/rest/v5
//
// It will provide a client and support objects that can be used to interact
// with Pivotal Tracker. This package does *not* handle any form of
// authentication, so the burden of security is put on the library user.
package tracker

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . HTTPClient

// HTTPClient is based off of http.Client.
type HTTPClient interface {
	Do(req *http.Request) (resp *http.Response, err error)
}

// Client is the Pivotal Tracker API Client. This client requires the base
// Pivotal Tracker API URL and a user's API Token. The user's API token will be
// used to make every request to the Pivotal Tracker API. This means that all
// modifications to Pivotal Tracker will be seen as the given user.
type Client struct {
	// APIURL is the base Pivotal Tracker API URL.
	APIURL *url.URL

	// APIToken is the user's API Token.
	APIToken APIToken

	// HTTPClient is a basic http client that can handle a Do request.
	HTTPClient HTTPClient
}

// NewClient will return a new Client with HTTPClient set to
// http.DefaultClient.
func NewClient(apiURL string, apiToken APIToken) (Client, error) {
	url, err := url.Parse(apiURL)
	if err != nil {
		return Client{}, err
	}

	return Client{
		APIURL:     url,
		APIToken:   apiToken,
		HTTPClient: http.DefaultClient,
	}, nil
}

func (c Client) get(uri string) ([]byte, error) {
	URL, err := c.APIURL.Parse(uri)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("GET", URL.String(), nil)
	if err != nil {
		return nil, err
	}

	c.APIToken.AddToRequestHeader(request)
	response, err := c.HTTPClient.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	return ioutil.ReadAll(response.Body)
}
