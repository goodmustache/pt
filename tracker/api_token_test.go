package tracker_test

import (
	"net/http"

	. "github.com/goodmustache/pt/tracker"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("ApiToken", func() {
	Describe("Validate", func() {
		DescribeTable("when the API token is invalid",
			func(tokenValue string, err error) {
				token := APIToken(tokenValue)
				Expect(token.Validate()).To(Equal(err))
			},

			Entry("errors due to invalid characters", "f@fffQffffffffffffffffffff$fffff", ErrorInvalidAPIToken),
			Entry("errors due to invalid size", "f", ErrorInvalidAPIToken),
		)
	})

	Describe("AddToRequestHeader", func() {
		It("adds the token as an X-TrackerToken header", func() {
			request, err := http.NewRequest("GET", "dun matter", nil)
			Expect(err).ToNot(HaveOccurred())

			apiToken := "foobar"
			APIToken(apiToken).AddToRequestHeader(request)
			Expect(request.Header.Get("X-TrackerToken")).To(Equal(apiToken))
		})
	})
})
