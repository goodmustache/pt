package tracker_test

import (
	"bytes"
	"log"
	"testing"

	. "github.com/goodmustache/pt/tracker"
	"github.com/goodmustache/pt/tracker/internal/internalfakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/ghttp"
)

func TestTracker(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Tracker API Suite")
}

var server *Server

var _ = SynchronizedBeforeSuite(func() []byte {
	return []byte{}
}, func(data []byte) {
	server = NewTLSServer()

	// Suppresses ginkgo server logs
	server.HTTPTestServer.Config.ErrorLog = log.New(&bytes.Buffer{}, "", 0)
})

var _ = SynchronizedAfterSuite(func() {
	server.Close()
}, func() {})

var _ = BeforeEach(func() {
	server.Reset()
})

func NewTestClient() (*Client, *internalfakes.FakeClock) {
	clock := new(internalfakes.FakeClock)
	return TestClient(
		server.URL(),
		clock,
		Config{
			AppName:           "Test Client",
			AppVersion:        "Unit Test",
			SkipSSLValidation: true,
		},
	), clock
}
