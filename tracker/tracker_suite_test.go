package tracker_test

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/goodmustache/pt/tracker"
	"github.com/goodmustache/pt/tracker/trackerfakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestTracker(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Tracker Suite")
}

const trackerURL = "https://fake.tracker.com"
const apiToken = "FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF"

var httpClient *trackerfakes.FakeHTTPClient
var client tracker.Client

var _ = BeforeEach(func() {
	var err error
	client, err = tracker.NewClient(trackerURL, apiToken)
	Expect(err).ToNot(HaveOccurred())

	httpClient = new(trackerfakes.FakeHTTPClient)
	client.HTTPClient = httpClient
})

func JSONResponse(body string) *http.Response {
	return &http.Response{
		Body: ioutil.NopCloser(bytes.NewBufferString(body)),
	}
}
