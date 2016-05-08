package actions_test

import (
	"time"

	. "github.com/goodmustache/pt/actions"
	"github.com/goodmustache/pt/config"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("User", func() {
	Describe("GetUser", func() {
		Context("when there are no issues reading the config", func() {
			Context("user exists", func() {
				DescribeTable("returns the user",
					func(alias string, username string, expectedID uint64) {
						conf := config.Config{
							CurrentUserID: 4,
							Users: []config.User{
								{ID: 2, Username: "agaitonde", Alias: "ag"},
								{ID: 3, Username: "hventure", Alias: "hv"},
								{ID: 4, Username: "dventure", Alias: "dv"},
							},
						}

						err := WriteConfig(conf)
						Expect(err).ToNot(HaveOccurred())

						user, err := GetUser(alias, username)
						Expect(err).ToNot(HaveOccurred())
						Expect(user.ID).To(Equal(expectedID))
					},

					Entry("for alias", "ag", "", uint64(2)),
					Entry("for username", "", "hventure", uint64(3)),
					Entry("for default", "", "", uint64(4)),
				)
			})

			Context("user exists", func() {
				DescribeTable("returns error",
					func(alias string, username string, expectedError error) {
						err := WriteConfig(config.Config{})
						Expect(err).ToNot(HaveOccurred())

						_, err = GetUser(alias, username)
						Expect(err).To(Equal(expectedError))
					},

					Entry("ErrUnableToFindAlias", "ga", "", ErrUnableToFindAlias),
					Entry("ErrUnableToFindUsername", "", "bsamson", ErrUnableToFindUsername),
					Entry("ErrNoCurrentUserSet", "", "", ErrNoCurrentUserSet),
					Entry("ErrBothAliasAndUsernameProvided", "ag", "hventure", ErrBothAliasAndUsernameProvided),
				)
			})
		})

		Context("when the config file does not exist", func() {
			It("returns an ErrNoCurrentUserSet", func() {
				_, err := GetUser("ag", "")
				Expect(err).To(Equal(ErrNoCurrentUserSet))
			})
		})
	})

	Describe("RemoveUser", func() {
		var userToRemove ConfigUser
		var userToKeep config.User

		BeforeEach(func() {
			userToKeep = config.User{ID: 2, Username: "agaitonde", Alias: "ag"}
			userToRemove = ConfigUser{ID: 3, Username: "hventure", Alias: "hv"}
			conf := config.Config{
				CurrentUserID:      userToKeep.ID,
				CurrentUserSetTime: time.Date(2014, 4, 14, 17, 6, 0, 0, time.UTC),
				Users: []config.User{
					userToKeep,
					config.User(userToRemove),
				},
			}

			err := WriteConfig(conf)
			Expect(err).ToNot(HaveOccurred())
		})

		It("removes user from config", func() {
			err := RemoveUser(userToRemove)
			Expect(err).ToNot(HaveOccurred())

			readConf, err := ReadConfig()
			Expect(err).ToNot(HaveOccurred())
			Expect(readConf).To(Equal(config.Config{
				CurrentUserID:      userToKeep.ID,
				CurrentUserSetTime: time.Date(2014, 4, 14, 17, 6, 0, 0, time.UTC),
				Users:              []config.User{userToKeep},
			}))
		})
	})
})
