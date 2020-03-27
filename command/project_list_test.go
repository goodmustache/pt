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

var _ = Describe("ProjectList", func() {
	var (
		cmd     ProjectList
		execErr error

		fakeActor *commandfakes.FakeProjectListActor
		testUI    *ui.UI
		out       *Buffer
	)

	BeforeEach(func() {
		fakeActor = new(commandfakes.FakeProjectListActor)
		testUI, _, out, _ = NewTestUI()

		cmd = ProjectList{
			Actor: fakeActor,
			UI:    testUI,
		}
	})

	JustBeforeEach(func() {
		execErr = cmd.Execute(nil)
	})

	When("there are projects", func() {
		BeforeEach(func() {
			fakeActor.ProjectsReturns([]actor.Project{
				{
					ID:          51,
					Name:        "Jeff",
					Description: "some description",
					Public:      false,
				},
				{
					ID:          11,
					Name:        "Anand",
					Description: "another description",
					Public:      true,
				},
			}, nil)
		})

		It("displays the projects in a table", func() {
			Expect(execErr).ToNot(HaveOccurred())

			Expect(testUI.STDOUT).To(Say(`ID\s+VISIBILITY\s+NAME\s+DESCRIPTION`))
			Expect(testUI.STDOUT).To(Say(`51\s+private\s+Jeff\s+some description`))
			Expect(testUI.STDOUT).To(Say(`11\s+public\s+Anand\s+another description`))
		})
	})

	When("listing projects errors", func() {
		BeforeEach(func() {
			fakeActor.ProjectsReturns(nil, errors.New("potato"))
		})

		It("displays the projects in a table", func() {
			Expect(execErr).To(MatchError("potato"))
			Expect(out.Contents()).To(BeEmpty())
		})
	})
})
