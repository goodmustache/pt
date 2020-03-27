package command_test

import (
	"errors"

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
		out        *Buffer
	)

	BeforeEach(func() {
		fakeConfig = new(commandfakes.FakeConfig)
		testUI, _, out, _ = NewTestUI()

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
					APIToken: "some-token",
					Username: "A1",
				},
				{
					Email:    "test2@email.com",
					ID:       11,
					Name:     "Anand",
					APIToken: "some-other-token",
					Username: "A2",
				},
			}, nil)
		})

		It("displays the users in a table", func() {
			Expect(execErr).ToNot(HaveOccurred())

			Expect(testUI.STDOUT).To(Say(`USER ID\s+USERNAME\s+NAME\s+EMAIL`))
			Expect(testUI.STDOUT).To(Say(`51\s+A1\s+Jeff\s+test1@email\.com`))
			Expect(testUI.STDOUT).To(Say(`11\s+A2\s+Anand\s+test2@email\.com`))
		})
	})

	When("getting users errors", func() {
		BeforeEach(func() {
			fakeConfig.GetUsersReturns(nil, errors.New("oh noes"))
		})

		It("displays the users in a table", func() {
			Expect(execErr).To(MatchError("oh noes"))
			Expect(out.Contents()).To(BeEmpty())
		})
	})
})
