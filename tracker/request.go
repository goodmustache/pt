package tracker

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/goodmustache/pt/tracker/internal"
)

// Params map path keys to values. For example, if your route has the path
// pattern:
//
//	/person/:person_id/pets/:pet_type
//
// Then a correct Params map would lool like:
//
//	router.Params{
//	  "person_id": "123",
//	  "pet_type": "cats",
//	}
type Params map[string]string

// requestOptions contains all the options to create an HTTP request.
type requestOptions struct {
	// URIParams are the list URI route parameters
	URIParams Params

	// Query is a list of HTTP query parameters. Query will overwrite any
	// existing query string in the URI. If you want to preserve the query
	// string in URI make sure Query is nil.
	// Query []Query

	// RequestName is the name of the request (see routes)
	RequestName string

	// Method is the HTTP method.
	Method string
	// URL is the request path.
	URL string

	// EncodeData is the object to json marshal in the request body.
	EncodeData interface{}
}

func (client *Client) MakeRequest(options requestOptions, returnResponse interface{}) error {
	var body io.ReadSeeker
	if options.EncodeData != nil {
		raw, err := json.Marshal(options.EncodeData)
		if err != nil {
			return err
		}

		body = bytes.NewReader(raw)
	}

	request, err := client.newHTTPRequest(options, body)
	if err != nil {
		return err
	}

	response := new(internal.Response)
	if returnResponse != nil {
		response.DecodeJSONResponseInto = returnResponse
	}

	return client.connection.Make(request, response)
}

// newHTTPRequest returns a constructed HTTP.Request with some defaults.
// Defaults are applied when Request options are not filled in.
func (client *Client) newHTTPRequest(passedRequest requestOptions, body io.ReadSeeker) (*internal.Request, error) {
	var request *http.Request
	var err error
	if passedRequest.URL != "" {
		request, err = http.NewRequest(
			passedRequest.Method,
			passedRequest.URL,
			body,
		)
	} else {
		request, err = client.router.CreateRequest(
			passedRequest.RequestName,
			map[string]string(passedRequest.URIParams),
			body,
		)
	}
	if err != nil {
		return nil, err
	}
	// if passedRequest.Query != nil {
	// 	request.URL.RawQuery = FormatQueryParameters(passedRequest.Query).Encode()
	// }

	request.Header = http.Header{}
	request.Header.Set("Accept", "application/json")
	request.Header.Set("User-Agent", client.userAgent)

	if body != nil {
		request.Header.Set("Content-Type", "application/json")
	}

	return internal.NewRequest(request, body), nil
}
