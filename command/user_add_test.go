package command_test

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"

	"github.com/goodmustache/pt/actor"
	. "github.com/goodmustache/pt/command"
	"github.com/goodmustache/pt/command/commandfakes"
	"github.com/goodmustache/pt/ui"
)

var _ = Describe("UserAdd", func() {
	var (
		cmd     UserAdd
		execErr error

		fakeActor  *commandfakes.FakeUserAddActor
		fakeConfig *commandfakes.FakeConfig
		testUI     *ui.UI

		token string
	)

	BeforeEach(func() {
		fakeActor = new(commandfakes.FakeUserAddActor)
		fakeConfig = new(commandfakes.FakeConfig)
		testUI = ui.NewTestUI(nil, NewBuffer(), nil)

		token = "some-token"

		cmd = UserAdd{
			APIToken: token,
			Actor:    fakeActor,
			Config:   fakeConfig,
			UI:       testUI,
		}
	})

	JustBeforeEach(func() {
		execErr = cmd.Execute(nil)
	})

	When("there are users", func() {
		When("when the user gets added properly", func() {
			BeforeEach(func() {
				fakeActor.AddUserReturns(actor.User{
					Email:    "test2@email.com",
					ID:       11,
					Name:     "Anand",
					Username: "A2",
				}, nil)
			})

			It("doesn't error", func() {
				Expect(execErr).ToNot(HaveOccurred())
			})

			It("adds the user to config", func() {
				Expect(fakeActor.AddUserCallCount()).To(Equal(1))
			})

			It("displays the user to add", func() {
				Expect(testUI.STDOUT).To(Say("USER ID\\s+USERNAME\\s+NAME\\s+EMAIL"))
				Expect(testUI.STDOUT).To(Say("11\\s+A2\\s+Anand\\s+test2@email\\.com"))
			})
		})

		When("when the user gets added properly", func() {
			BeforeEach(func() {
				fakeActor.AddUserReturns(actor.User{}, errors.New("oh noes"))
			})

			It("doesn't error", func() {
				Expect(execErr).To(MatchError("oh noes"))

				Expect(testUI.STDOUT).ToNot(Say("USERNAME\\s+NAME\\s+EMAIL"))
			})
		})
	})
})
