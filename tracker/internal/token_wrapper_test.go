package internal_test

import (
	"errors"
	"net/http"

	. "github.com/goodmustache/pt/tracker/internal"
	"github.com/goodmustache/pt/tracker/internal/internalfakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Token Wrapper", func() {
	var (
		fakeConnection *internalfakes.FakeConnection
		wrapper        *TokenWrapper

		token string
	)

	BeforeEach(func() {
		token = "some-token"
		wrapper = NewTokenWrapper(token)

		fakeConnection = new(internalfakes.FakeConnection)
		wrapper.Wrap(fakeConnection)
	})

	Describe("Make", func() {
		var (
			request  *Request
			response *Response

			executeErr error
		)

		BeforeEach(func() {
			req, err := http.NewRequest("", "", nil)
			Expect(err).ToNot(HaveOccurred())

			request = &Request{
				Request: req,
			}
			response = new(Response)
		})

		JustBeforeEach(func() {
			executeErr = wrapper.Make(request, response)
		})

		It("adds the token to the request header", func() {
			Expect(fakeConnection.MakeCallCount()).To(Equal(1))
			passedRequest, _ := fakeConnection.MakeArgsForCall(0)
			Expect(passedRequest.Header.Get("X-TrackerToken")).To(Equal(token))
		})

		It("passes through the response body", func() {
			Expect(fakeConnection.MakeCallCount()).To(Equal(1))
		})

		When("there's an error from the wrapped connection's make", func() {
			BeforeEach(func() {
				fakeConnection.MakeReturns(errors.New("internal error"))
			})

			It("returns the error", func() {
				Expect(executeErr).To(MatchError("internal error"))
			})
		})
	})
})
