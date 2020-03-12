package tracker_test

import (
	"fmt"
	"net/http"
	"runtime"

	. "github.com/goodmustache/pt/tracker"
	"github.com/goodmustache/pt/tracker/internal/internalfakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/ghttp"
)

var _ = Describe("Client", func() {
	var (
		client *Client
	)

	BeforeEach(func() {
		client, _ = NewTestClient()
	})

	Describe("WrapConnection", func() {
		var fakeConnectionWrapper *internalfakes.FakeConnectionWrapper

		BeforeEach(func() {
			fakeConnectionWrapper = new(internalfakes.FakeConnectionWrapper)
			fakeConnectionWrapper.WrapReturns(fakeConnectionWrapper)
		})

		It("wraps the existing connection in the provided wrapper", func() {
			client.WrapConnection(fakeConnectionWrapper)
			Expect(fakeConnectionWrapper.WrapCallCount()).To(Equal(1))

			_, err := client.Me()
			Expect(err).ToNot(HaveOccurred())
			Expect(fakeConnectionWrapper.MakeCallCount()).To(Equal(1))
		})
	})

	Describe("User Agent", func() {
		BeforeEach(func() {
			expectedUserAgent := fmt.Sprintf("Test Client/Unit Test (%s; %s %s)", runtime.Version(), runtime.GOARCH, runtime.GOOS)
			server.AppendHandlers(
				CombineHandlers(
					VerifyHeaderKV("User-Agent", expectedUserAgent),
					RespondWith(http.StatusOK, "{}"),
				),
			)
		})

		It("adds a user agent header", func() {
			_, err := client.Me()
			Expect(err).ToNot(HaveOccurred())
			Expect(server.ReceivedRequests()).To(HaveLen(1))
		})
	})
})
