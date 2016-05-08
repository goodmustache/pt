package commands_test

import (
	"github.com/goodmustache/pt/actions"
	"github.com/goodmustache/pt/config"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
	. "github.com/onsi/gomega/gexec"
)

var _ = Describe("User", func() {
	Describe("Help", func() {
		It("displays help for user", func() {
			session := runCommand("user", "-h")

			Eventually(session).ShouldNot(Exit(0))
			Expect(session.Out).To(Say("user"))
			Expect(session.Out).To(Say("--alias"))
			Expect(session.Out).To(Say("--username"))
		})
	})

	Context("when the current user exists", func() {
		var (
			conf  config.Config
			user2 config.User
			user3 config.User
			user4 config.User
		)

		BeforeEach(func() {
			user2 = config.User{ID: 2, Name: "Anand Gaitonde", Username: "agaitonde", Alias: "ag"}
			user3 = config.User{ID: 3, Name: "Hank Venture", Username: "hventure", Alias: "hv"}
			user4 = config.User{ID: 4, Name: "Dean Venture", Username: "dventure", Alias: "dv"}

			conf = config.Config{
				CurrentUserID: user3.ID,
				Users: []config.User{
					user2,
					user3,
					user4,
				},
			}

			err := actions.WriteConfig(conf)
			Expect(err).ToNot(HaveOccurred())
		})

		It("returns the current user", func() {
			session := runCommand("user")

			Eventually(session.Out).Should(Say("Name: %s", user3.Name))
			Eventually(session.Out).Should(Say("Username: %s", user3.Username))

			Eventually(session).Should(Exit(0))
		})

		It("returns the user set with -a", func() {
			session := runCommand("user", "-a", user2.Alias)

			Eventually(session.Out).Should(Say("Name: %s", user2.Name))
			Eventually(session.Out).Should(Say("Username: %s", user2.Username))

			Eventually(session).Should(Exit(0))
		})

		It("returns the user set with -u", func() {
			session := runCommand("user", "-u", user4.Username)

			Eventually(session.Out).Should(Say("Name: %s", user4.Name))
			Eventually(session.Out).Should(Say("Username: %s", user4.Username))

			Eventually(session).Should(Exit(0))
		})
	})
})
