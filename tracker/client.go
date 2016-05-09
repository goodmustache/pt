package tracker

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
)

//go:generate counterfeiter . HTTPClient

type HTTPClient interface {
	Do(req *http.Request) (resp *http.Response, err error)
}

type Client struct {
	APIURL     *url.URL
	APIToken   APIToken
	HTTPClient HTTPClient
}

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
	request, err := http.NewRequest("GET", path.Join(c.APIURL.Path, uri), nil)
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
