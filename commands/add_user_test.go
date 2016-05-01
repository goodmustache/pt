package commands_test

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"path"
	"time"

	. "github.com/goodmustache/pt/commands"
	"github.com/goodmustache/pt/config"
	"github.com/goodmustache/pt/tracker"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
	. "github.com/onsi/gomega/gexec"
	. "github.com/onsi/gomega/ghttp"
)

var _ = Describe("AddUser", func() {
	var apiToken string

	BeforeEach(func() {
		apiToken = "FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF"
	})

	Describe("Help", func() {
		It("displays help for add-user", func() {
			session := runCommand("add-user", "-h")

			Eventually(session).ShouldNot(Exit(0))
			Expect(session.Out).To(Say("add-user"))
			Expect(session.Out).To(Say("api-token"))
		})
	})

	Describe("API Token", func() {
		var tokenInfo tracker.TokenInfomation
		var expectedConfig config.Config

		Context("when config doesn't exist", func() {
			BeforeEach(func() {
				tokenInfo = tracker.TokenInfomation{
					APIToken: apiToken,
					ID:       42,
					Name:     "Anand Gaitonde",
					Username: "agaitonde",
				}

				expectedConfig = config.Config{
					CurrentUserID: tokenInfo.ID,
					Users: []config.User{
						{
							APIToken: tokenInfo.APIToken,
							ID:       tokenInfo.ID,
							Name:     tokenInfo.Name,
							Username: tokenInfo.Username,
						},
					},
				}

				server.AppendHandlers(
					CombineHandlers(
						VerifyRequest("GET", "/me"),
						RespondWithJSONEncoded(http.StatusOK, tokenInfo),
					),
				)
			})

			It("prompts for api token if not passed", func() {
				session, stdin := runCommandWithInput("add-user")

				Eventually(session.Out).Should(Say(AddUserInstructions))

				Eventually(session.Out).Should(Say("API Token:"))
				inputValue(apiToken, stdin)

				Eventually(session.Out).Should(Say("Added User! Setting %s \\(%s\\) to be the current user.", tokenInfo.Name, tokenInfo.Username))

				Eventually(session).Should(Exit(0))

				rawConfig, err := ioutil.ReadFile(path.Join(userHomeDir(), ".config", "pt", "config.json"))
				Expect(err).ToNot(HaveOccurred())

				var config config.Config
				err = json.Unmarshal(rawConfig, &config)
				Expect(err).ToNot(HaveOccurred())

				config.CurrentUserSetTime = time.Time{}
				Expect(config).To(Equal(expectedConfig))
			})

			It("does not prompt if api token is not passed", func() {
				session := runCommand("add-user", "--api-token", apiToken)

				Eventually(session.Out).ShouldNot(Say(AddUserInstructions))
				Eventually(session.Out).ShouldNot(Say("API Token:"))

				Eventually(session.Out).Should(Say("Added User! Setting %s \\(%s\\) to be the current user.", tokenInfo.Name, tokenInfo.Username))
				Eventually(session).Should(Exit(0))

				rawConfig, err := ioutil.ReadFile(path.Join(userHomeDir(), ".config", "pt", "config.json"))
				Expect(err).ToNot(HaveOccurred())

				var config config.Config
				err = json.Unmarshal(rawConfig, &config)
				Expect(err).ToNot(HaveOccurred())

				config.CurrentUserSetTime = time.Time{}
				Expect(config).To(Equal(expectedConfig))
			})
		})
	})
})
