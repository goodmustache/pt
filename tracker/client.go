package tracker

import (
	"io/ioutil"
	"net/http"
)

type TrackerClient interface {
	TokenInfo() (TokenInfomation, error)
}

type client struct {
	APIURL     string
	APIToken   string
	httpClient *http.Client
}

func NewTrackerClient(apiURL string, apiToken string) TrackerClient {
	return &client{
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

	request.Header.Add("X-TrackerToken", c.APIToken)
	response, err := c.httpClient.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	return ioutil.ReadAll(response.Body)
}
