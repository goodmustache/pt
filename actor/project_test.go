package actor_test

import (
	"errors"

	. "github.com/goodmustache/pt/actor"
	"github.com/goodmustache/pt/actor/actorfakes"
	"github.com/goodmustache/pt/tracker"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Project", func() {
	var project Project

	DescribeTable("Visibility",
		func(trackerValue bool, expectedReturn string) {
			project.Public = trackerValue
			Expect(project.Visibility()).To(Equal(expectedReturn))
		},
		Entry("when public is true, returns public", true, "public"),
		Entry("when public is false, returns private", false, "private"),
	)
})

var _ = Describe("Project Actions", func() {
	var (
		actor   *Main
		fakeAPI *actorfakes.FakeTrackerClient
	)

	BeforeEach(func() {
		actor, fakeAPI, _ = NewTestActor()
	})

	Describe("Projects", func() {
		var (
			projects   []Project
			executeErr error
		)

		JustBeforeEach(func() {
			projects, executeErr = actor.Projects()
		})

		When("requesting user information is successful", func() {
			BeforeEach(func() {
				returnedProjects := []tracker.Project{
					{ID: 1, Name: "Foobar"},
					{ID: 2, Name: "Azure"},
					{ID: 3, Name: "fafofo"},
					{ID: 4, Name: "bacon"},
				}
				fakeAPI.ProjectsReturns(returnedProjects, nil)
			})

			It("does not error", func() {
				Expect(executeErr).ToNot(HaveOccurred())
			})

			It("sorts the projects projects by name (case insensitive alphanumeric)", func() {
				Expect(projects).To(Equal([]Project{
					{ID: 2, Name: "Azure"},
					{ID: 4, Name: "bacon"},
					{ID: 3, Name: "fafofo"},
					{ID: 1, Name: "Foobar"},
				}))
			})
		})

		When("requesting projects errors", func() {
			BeforeEach(func() {
				fakeAPI.ProjectsReturns(nil, errors.New("me error"))
			})

			It("returns the requested user", func() {
				Expect(executeErr).To(MatchError("me error"))
			})
		})
	})
})
