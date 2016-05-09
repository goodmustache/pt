package commands_test

import (
	"net/http"

	"github.com/goodmustache/pt/actions"
	. "github.com/goodmustache/pt/commands"
	"github.com/goodmustache/pt/config"
	"github.com/goodmustache/pt/tracker"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
	. "github.com/onsi/gomega/gexec"
	. "github.com/onsi/gomega/ghttp"
)

var _ = Describe("Add User", func() {
	const apiToken = "FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF"

	Describe("Help", func() {
		It("displays help for add-user", func() {
			session := runCommand("add-user", "-h")

			Eventually(session).ShouldNot(Exit(0))
			Expect(session.Err).To(Say("add-user"))
			Expect(session.Err).To(Say("--api-token"))
			Expect(session.Err).To(Say("--alias"))
		})
	})

	DescribeTable("input prompts and output",
		func(inputApiToken func() *Session) {
			tokenInfo := tracker.TokenInformation{
				APIToken: apiToken,
				ID:       42,
				Name:     "Anand Gaitonde",
				Username: "agaitonde",
			}

			server.AppendHandlers(
				CombineHandlers(
					VerifyRequest("GET", "/me"),
					VerifyHeader(http.Header{"X-TrackerToken": []string{apiToken}}),
					RespondWithJSONEncoded(http.StatusOK, tokenInfo),
				),
			)

			session := inputApiToken()

			Eventually(session.Out).Should(Say("Added User! Setting %s \\(%s\\) to be the current user.", tokenInfo.Name, tokenInfo.Username))
			Eventually(session).Should(Exit(0))
		},

		Entry("prompts for api token when not passed", func() *Session {
			session, stdin := runCommandWithInput("add-user")
			defer stdin.Close()

			Eventually(session.Out).Should(Say(AddUserInstructions))

			Eventually(session.Out).Should(Say("API Token:"))
			inputValue(apiToken, stdin)
			return session
		}),

		Entry("does not prompt when api token is passed", func() *Session {
			session := runCommand("add-user", "--api-token", apiToken)

			Eventually(session.Out).ShouldNot(Say(AddUserInstructions))
			Eventually(session.Out).ShouldNot(Say("API Token:"))
			return session
		}),
	)

	Context("writing the config", func() {
		var tokenInfo tracker.TokenInformation

		BeforeEach(func() {
			tokenInfo = tracker.TokenInformation{
				APIToken: apiToken,
				ID:       42,
				Name:     "Anand Gaitonde",
				Username: "agaitonde",
			}

			server.AppendHandlers(
				CombineHandlers(
					VerifyRequest("GET", "/me"),
					RespondWithJSONEncoded(http.StatusOK, tokenInfo),
				),
			)
		})

		It("appends the new user to the config", func() {
			alias := "ag"
			session := runCommand("add-user", "--api-token", apiToken, "-a", alias)

			Eventually(session).Should(Exit(0))

			readConf, err := actions.ReadConfig()
			Expect(err).ToNot(HaveOccurred())

			users := readConf.Users
			Expect(users).To(HaveLen(1))
			Expect(users[0]).To(Equal(config.User{
				ID:       tokenInfo.ID,
				APIToken: tokenInfo.APIToken,
				Name:     tokenInfo.Name,
				Username: tokenInfo.Username,
				Alias:    alias,
			}))
		})
	})
})
