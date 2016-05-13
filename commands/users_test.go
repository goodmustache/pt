package commands_test

import (
	"github.com/goodmustache/pt/actions"
	. "github.com/goodmustache/pt/commands"
	"github.com/goodmustache/pt/config"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
	. "github.com/onsi/gomega/gexec"
)

var _ = Describe("Users", func() {
	Describe("Help", func() {
		It("displays help for users", func() {
			session := runCommand("user", "-h")

			Eventually(session).ShouldNot(Exit(0))
			Expect(session.Err).To(Say("user"))
		})
	})

	Describe("listing users", func() {
		var session *Session

		JustBeforeEach(func() {
			session = runCommand("users")
		})

		Context("config exists", func() {
			Context("when there are users", func() {
				var conf config.Config

				BeforeEach(func() {
					conf = config.Config{
						Users: []config.User{
							{
								ID:       4,
								Name:     "Anand Gaitonde",
								Username: "agaitonde",
								Alias:    "ag",
							},
							{
								ID:       5,
								Name:     "Hank Venture",
								Username: "hventure",
								Alias:    "hv",
							},
							{
								ID:       6,
								Name:     "Brock Samson",
								Username: "bsamson",
							},
						},
					}

					err := actions.WriteConfig(conf)
					Expect(err).ToNot(HaveOccurred())
				})

				It("informs users that no users have been added", func() {
					Expect(session.Out).To(Say("Name\\s+Username\\s+Alias"))

					for _, user := range conf.Users {
						Expect(session.Out).To(Say("%s\\s+%s\\s+%s", user.Name, user.Username, user.Alias))
					}
					Expect(session).To(Exit(0))
				})
			})

			Context("when there are no users", func() {
				BeforeEach(func() {
					err := actions.WriteConfig(config.Config{})
					Expect(err).ToNot(HaveOccurred())
				})

				It("informs users that no users have been added", func() {
					Expect(session.Err).To(Say(ErrNoUsers.Error()))
					Expect(session).ToNot(Exit(0))
				})
			})
		})

		Context("config does not exist", func() {
			It("informs users that no users have been added", func() {
				Expect(session.Err).To(Say(ErrNoUsers.Error()))
				Expect(session).ToNot(Exit(0))
			})
		})
	})
})
