package actor_test

import (
	"errors"

	. "github.com/goodmustache/pt/actor"
	"github.com/goodmustache/pt/actor/actorfakes"
	"github.com/goodmustache/pt/config"
	"github.com/goodmustache/pt/tracker"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("User Actions", func() {
	var (
		actor      *Main
		fakeAPI    *actorfakes.FakeTrackerClient
		fakeConfig *actorfakes.FakeConfig
	)

	BeforeEach(func() {
		actor, fakeAPI, fakeConfig = NewTestActor()
	})

	Describe("AddUser", func() {
		var (
			user       User
			executeErr error
		)

		JustBeforeEach(func() {
			user, executeErr = actor.AddUser()
		})

		When("requesting user information is successful", func() {
			var returnedUser tracker.Me

			BeforeEach(func() {
				returnedUser = tracker.Me{
					ID:   42,
					Name: "Anand",
				}
				fakeAPI.MeReturns(returnedUser, nil)
			})

			It("returns the requested user", func() {
				Expect(executeErr).ToNot(HaveOccurred())
				Expect(user).To(Equal(User(returnedUser)))
			})

			It("adds the user to the local config", func() {
				Expect(fakeConfig.AddUserCallCount()).To(Equal(1))
				Expect(fakeConfig.AddUserArgsForCall(0)).To(Equal(
					config.User{
						ID:   42,
						Name: "Anand",
					},
				))
			})

			When("adding the user errors", func() {
				BeforeEach(func() {
					fakeConfig.AddUserReturns(errors.New("add error"))
				})

				It("returns the requested user", func() {
					Expect(executeErr).To(MatchError("add error"))
				})
			})
		})

		When("requesting user information errors", func() {
			BeforeEach(func() {
				fakeAPI.MeReturns(tracker.Me{}, errors.New("me error"))
			})

			It("returns the requested user", func() {
				Expect(executeErr).To(MatchError("me error"))
			})
		})
	})
})
