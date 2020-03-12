package internal_test

import (
	"fmt"
	"net/http"

	. "github.com/goodmustache/pt/tracker/internal"
	"github.com/goodmustache/pt/tracker/terror"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/ghttp"
)

type DummyResponse struct {
	Val1 string      `json:"val1"`
	Val2 int         `json:"val2"`
	Val3 interface{} `json:"val3,omitempty"`
}

var _ = Describe("Tracker Connection", func() {
	var connection Connection

	BeforeEach(func() {
		connection = NewTrackerConnection(ConnectionConfig{SkipSSLValidation: true})
	})

	Describe("Make", func() {
		Describe("Data Unmarshalling", func() {
			var request *Request

			BeforeEach(func() {
				response := `{
					"val1":"2.59.0",
					"val2":2,
					"val3":1111111111111111111
				}`
				server.AppendHandlers(
					CombineHandlers(
						VerifyRequest(http.MethodGet, "/v2/foo", ""),
						RespondWith(http.StatusOK, response),
					),
				)

				req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/v2/foo", server.URL()), nil)
				Expect(err).ToNot(HaveOccurred())
				request = &Request{Request: req}
			})

			When("passed a response with a result set", func() {
				It("unmarshals the data into a struct", func() {
					var body DummyResponse
					response := Response{
						DecodeJSONResponseInto: &body,
					}

					err := connection.Make(request, &response)
					Expect(err).NotTo(HaveOccurred())

					Expect(body.Val1).To(Equal("2.59.0"))
					Expect(body.Val2).To(Equal(2))
				})

				It("keeps numbers unmarshalled to interfaces as interfaces", func() {
					var body DummyResponse
					response := Response{
						DecodeJSONResponseInto: &body,
					}

					err := connection.Make(request, &response)
					Expect(err).NotTo(HaveOccurred())
					Expect(fmt.Sprint(body.Val3)).To(Equal("1111111111111111111"))
				})
			})

			When("passed an empty response", func() {
				It("skips the unmarshalling step", func() {
					var response Response
					err := connection.Make(request, &response)
					Expect(err).NotTo(HaveOccurred())
					Expect(response.DecodeJSONResponseInto).To(BeNil())
				})
			})
		})

		Describe("HTTP Response", func() {
			var request *Request

			BeforeEach(func() {
				response := `{}`
				server.AppendHandlers(
					CombineHandlers(
						VerifyRequest(http.MethodGet, "/v2/foo", ""),
						RespondWith(http.StatusOK, response),
					),
				)

				req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/v2/foo", server.URL()), nil)
				Expect(err).ToNot(HaveOccurred())
				request = &Request{Request: req}
			})

			It("returns the status", func() {
				response := Response{}

				err := connection.Make(request, &response)
				Expect(err).NotTo(HaveOccurred())

				Expect(response.HTTPResponse.Status).To(Equal("200 OK"))
			})
		})

		Describe("Response Headers", func() {
			Describe("Location", func() {
				BeforeEach(func() {
					server.AppendHandlers(
						CombineHandlers(
							VerifyRequest(http.MethodGet, "/v2/foo"),
							RespondWith(http.StatusAccepted, "{}", http.Header{"Location": {"/v2/some-location"}}),
						),
					)
				})

				It("returns the location in the ResourceLocationURL", func() {
					req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/v2/foo", server.URL()), nil)
					Expect(err).ToNot(HaveOccurred())
					request := &Request{Request: req}

					var response Response
					err = connection.Make(request, &response)
					Expect(err).NotTo(HaveOccurred())

					Expect(server.ReceivedRequests()).To(HaveLen(1))
					Expect(response.ResourceLocationURL).To(Equal("/v2/some-location"))
				})
			})
		})

		Describe("Errors", func() {
			Describe("RawHTTPStatusError", func() {
				var ccResponse string
				BeforeEach(func() {
					ccResponse = `{
						"code": 90004,
						"description": "The service binding could not be found: some-guid",
						"error_code": "CF-ServiceBindingNotFound"
					}`

					server.AppendHandlers(
						CombineHandlers(
							VerifyRequest(http.MethodGet, "/v2/foo"),
							RespondWith(
								http.StatusNotFound,
								ccResponse,
								http.Header{
									"X-Request-Id": {"6e0b4379-f5f7-4b2b-56b0-9ab7e96eed95", "6e0b4379-f5f7-4b2b-56b0-9ab7e96eed95::7445d9db-c31e-410d-8dc5-9f79ec3fc26f"},
								}),
						),
					)
				})

				It("returns a CCRawResponse", func() {
					req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/v2/foo", server.URL()), nil)
					Expect(err).ToNot(HaveOccurred())
					request := &Request{Request: req}

					var response Response
					err = connection.Make(request, &response)
					Expect(err).To(MatchError(terror.RawHTTPStatusError{
						StatusCode:  http.StatusNotFound,
						RawResponse: []byte(ccResponse),
						RequestIDs:  []string{"6e0b4379-f5f7-4b2b-56b0-9ab7e96eed95", "6e0b4379-f5f7-4b2b-56b0-9ab7e96eed95::7445d9db-c31e-410d-8dc5-9f79ec3fc26f"},
					}))

					Expect(server.ReceivedRequests()).To(HaveLen(1))
				})
			})
		})
	})
})
