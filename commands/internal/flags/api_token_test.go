package flags_test

import (
	. "github.com/goodmustache/pt/commands/internal/flags"
	"github.com/goodmustache/pt/tracker"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ApiToken", func() {
	var apiToken APIToken
	BeforeEach(func() {
		apiToken = APIToken{}
	})

	Context("valid api tokens", func() {
		It("parses and stores", func() {
			tokenValue := "FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF"
			err := apiToken.UnmarshalFlag(tokenValue)
			Expect(err).ToNot(HaveOccurred())
			Expect(apiToken.Value).To(Equal(tracker.APIToken(tokenValue)))
		})
	})

	Context("invalid api tokens", func() {
		It("errors", func() {
			err := apiToken.UnmarshalFlag("")
			Expect(err).To(HaveOccurred())
		})
	})
})
