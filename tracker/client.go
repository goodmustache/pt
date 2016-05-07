package tracker

import (
	"io/ioutil"
	"net/http"
)

type client struct {
	APIURL     string
	APIToken   APIToken
	httpClient *http.Client
}

func NewClient(apiURL string, apiToken APIToken) client {
	return client{
		APIURL:     apiURL,
		APIToken:   apiToken,
		httpClient: http.DefaultClient,
	}
}

func (c client) get(uri string) ([]byte, error) {
	request, err := http.NewRequest("GET", c.APIURL+uri, nil)
	if err != nil {
		return nil, err
	}

	c.APIToken.AddToRequestHeader(request)
	response, err := c.httpClient.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	return ioutil.ReadAll(response.Body)
}
