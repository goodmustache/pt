package commands_test

import (
	"time"

	"github.com/goodmustache/pt/actions"
	"github.com/goodmustache/pt/config"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
	. "github.com/onsi/gomega/gexec"
)

var _ = Describe("Remove User", func() {
	Describe("Help", func() {
		It("displays help for remove-user", func() {
			session := runCommand("remove-user", "-h")

			Eventually(session).ShouldNot(Exit(0))
			Expect(session.Err).To(Say("remove-user"))
			Expect(session.Err).To(Say("--alias"))
			Expect(session.Err).To(Say("--username"))
			Expect(session.Err).To(Say("--with-malice"))
		})
	})

	var user1 = config.User{ID: 2, APIToken: "doesn't matter", Name: "Anand Gaitonde", Username: "agaitonde", Alias: "ag"}
	var user2 = config.User{ID: 3, APIToken: "doesn't matter", Name: "Hank Venture", Username: "hventure", Alias: "hv"}

	DescribeTable("input prompts and output",
		func(removeUserCmd func() (*Session, config.User)) {
			conf := config.Config{
				CurrentUserID:      user1.ID,
				CurrentUserSetTime: time.Now(),
				Users:              []config.User{user1, user2},
			}
			err := actions.WriteConfig(conf)
			Expect(err).ToNot(HaveOccurred())

			session, removedUser := removeUserCmd()

			Eventually(session.Out).Should(Say("User %s \\(%s\\) has been removed.", removedUser.Name, removedUser.Username))
			Eventually(session).Should(Exit(0))
		},

		Entry("prompts for removal of current user", func() (*Session, config.User) {
			session, stdin := runCommandWithInput("remove-user")
			defer stdin.Close()

			Eventually(session.Out).Should(Say("Remove %s \\(%s\\):", user1.Name, user1.Username))
			inputValue("yes", stdin)
			return session, user1
		}),

		Entry("prompts for removal of current user provided by alias", func() (*Session, config.User) {
			session, stdin := runCommandWithInput("remove-user", "-a", user1.Alias)
			defer stdin.Close()

			Eventually(session.Out).Should(Say("Remove %s \\(%s\\):", user1.Name, user1.Username))
			inputValue("yes", stdin)
			return session, user1
		}),

		Entry("prompts for removal of current user provided by username", func() (*Session, config.User) {
			session, stdin := runCommandWithInput("remove-user", "-u", user1.Username)
			defer stdin.Close()

			Eventually(session.Out).Should(Say("Remove %s \\(%s\\):", user1.Name, user1.Username))
			inputValue("yes", stdin)
			return session, user1
		}),

		Entry("prompts for removal of different user provided by username", func() (*Session, config.User) {
			session, stdin := runCommandWithInput("remove-user", "-u", user2.Username)
			defer stdin.Close()

			Eventually(session.Out).Should(Say("Remove %s \\(%s\\):", user2.Name, user2.Username))
			inputValue("yes", stdin)
			return session, user2
		}),

		Entry("doesn't prompts for removal of different user provided by username if with-malice flag is provided", func() (*Session, config.User) {
			session, stdin := runCommandWithInput("remove-user", "-u", user2.Username, "--with-malice")
			defer stdin.Close()

			Eventually(session.Out).ShouldNot(Say("Remove %s \\(%s\\):", user2.Name, user2.Username))
			inputValue("yes", stdin)
			return session, user2
		}),
	)

	It("removes user from config", func() {
		conf := config.Config{
			CurrentUserID:      user1.ID,
			CurrentUserSetTime: time.Now(),
			Users:              []config.User{user1, user2},
		}

		err := actions.WriteConfig(conf)
		Expect(err).ToNot(HaveOccurred())

		session := runCommand("remove-user", "-u", user2.Username, "--with-malice")

		Eventually(session.Out).Should(Say("User %s \\(%s\\) has been removed.", user2.Name, user2.Username))
		Eventually(session).Should(Exit(0))

		readConfig, err := actions.ReadConfig()
		Expect(err).ToNot(HaveOccurred())
		Expect(readConfig.Users).To(HaveLen(1))
		Expect(readConfig.Users[0].ID).To(Equal(user1.ID))
	})

	It("does not remove user when no is inputted", func() {
		conf := config.Config{
			CurrentUserID:      user1.ID,
			CurrentUserSetTime: time.Now(),
			Users:              []config.User{user1, user2},
		}

		err := actions.WriteConfig(conf)
		Expect(err).ToNot(HaveOccurred())

		session, stdin := runCommandWithInput("remove-user", "-u", user2.Username)
		defer stdin.Close()

		Eventually(session.Out).Should(Say("Remove %s \\(%s\\):", user2.Name, user2.Username))
		inputValue("no", stdin)

		Eventually(session.Out).ShouldNot(Say("User %s \\(%s\\) has been removed.", user2.Name, user2.Username))
		Eventually(session).ShouldNot(Exit(0))
	})
})
