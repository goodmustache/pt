package command_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"

	. "github.com/goodmustache/pt/command"
	"github.com/goodmustache/pt/command/commandfakes"
	"github.com/goodmustache/pt/config"
	"github.com/goodmustache/pt/ui"
)

var _ = Describe("UserList", func() {
	var (
		cmd     UserList
		execErr error

		fakeConfig *commandfakes.FakeConfig
		testUI     *ui.UI
	)

	BeforeEach(func() {
		fakeConfig = new(commandfakes.FakeConfig)
		testUI = ui.NewTestUI(nil, NewBuffer(), nil)

		cmd = UserList{
			Config: fakeConfig,
			UI:     testUI,
		}
	})

	JustBeforeEach(func() {
		execErr = cmd.Execute(nil)
	})

	When("there are users", func() {
		BeforeEach(func() {
			fakeConfig.GetUsersReturns([]config.User{
				{
					Email:    "test1@email.com",
					ID:       51,
					Name:     "Jeff",
					Token:    "some-token",
					Username: "A1",
				},
				{
					Email:    "test2@email.com",
					ID:       11,
					Name:     "Anand",
					Token:    "some-other-token",
					Username: "A2",
				},
			}, nil)
		})

		It("displays the users in a table", func() {
			Expect(execErr).ToNot(HaveOccurred())

			Expect(testUI.STDOUT).To(Say("USERNAME\\s+NAME\\s+EMAIL"))
			Expect(testUI.STDOUT).To(Say("A1\\s+Jeff\\s+test1@email\\.com"))
			Expect(testUI.STDOUT).To(Say("A2\\s+Anand\\s+test2@email\\.com"))
		})
	})
})
