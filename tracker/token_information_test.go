package tracker_test

import (
	"fmt"

	. "github.com/goodmustache/pt/tracker"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Token Information", func() {
	BeforeEach(func() {
		tokenInfo := fmt.Sprintf(`{
			"id": 42,
			"api_token": "%s",
			"name": "Anand Gaitonde",
			"username": "agaitonde"
		}`, apiToken)

		httpClient.DoReturns(JSONResponse(tokenInfo), nil)
	})

	It("returns token information", func() {
		tokenInfo, err := client.TokenInformation()
		Expect(err).ToNot(HaveOccurred())
		Expect(tokenInfo).To(Equal(TokenInformation{
			ID:       42,
			APIToken: apiToken,
			Name:     "Anand Gaitonde",
			Username: "agaitonde",
		}))

		Expect(httpClient.DoCallCount()).To(Equal(1))
		request := httpClient.DoArgsForCall(0)
		Expect(request.URL.RequestURI()).To(Equal("/me"))
		Expect(request.Header.Get("X-TrackerToken")).To(Equal(apiToken))
	})
})
