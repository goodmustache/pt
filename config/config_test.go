package config_test

import (
	"time"

	. "github.com/goodmustache/pt/config"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Config", func() {
	var rawConfig []byte
	var parsedConfig Config

	BeforeEach(func() {
		rawConfig = []byte(`{
		"current_user_id": 42,
		"current_user_set_time": "2016-04-26T07:37:47Z",
		"users": [
			{
				"id": 45,
				"username": "dventure",
				"name": "Dean Venture",
				"api_token": "11111111111111111111111111111111",
				"alias": "dv"
			},
			{
				"id": 46,
				"username": "hventure",
				"name": "Henry Allen 'Hank' Venture",
				"api_token": "22222222222222222222222222222222",
				"alias": "hv"
			},
			{
				"id": 47,
				"username": "bsamson",
				"name": "Brock Samson",
				"api_token": "FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF"
			}
		]
	}`)

		parsedConfig = Config{
			CurrentUserID:      42,
			CurrentUserSetTime: time.Date(2016, 4, 26, 7, 37, 47, 0, time.UTC),
			Users: []User{
				{
					ID:       45,
					Username: "dventure",
					Name:     "Dean Venture",
					APIToken: "11111111111111111111111111111111",
					Alias:    "dv",
				},
				{
					ID:       46,
					Username: "hventure",
					Name:     "Henry Allen 'Hank' Venture",
					APIToken: "22222222222222222222222222222222",
					Alias:    "hv",
				},
				{
					ID:       47,
					Username: "bsamson",
					Name:     "Brock Samson",
					APIToken: "FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF",
				},
			},
		}
	})

	Describe("AddUser", func() {
		DescribeTable("adds the user",
			func(alias string) {
				newUser := User{
					ID:       48,
					APIToken: "ffffffffffffffffffffffffffffffff",
					Name:     "Dermott Fictel",
					Username: "TheWolf",
					Alias:    alias,
				}
				err := parsedConfig.AddUser(newUser.ID, newUser.APIToken, newUser.Name, newUser.Username, newUser.Alias)
				Expect(err).ToNot(HaveOccurred())

				Expect(parsedConfig.Users).To(HaveLen(4))
				Expect(parsedConfig.Users[3]).To(Equal(newUser))
			},

			Entry("with an alias", "tw"),
			Entry("without an alias", ""),
		)

		It("overrides user if id already exists", func() {
			newUser := User{
				ID:       45,
				APIToken: "ffffffffffffffffffffffffffffffff",
				Name:     "Dermott Fictel",
				Username: "TheWolf",
				Alias:    "tw",
			}

			err := parsedConfig.AddUser(newUser.ID, newUser.APIToken, newUser.Name, newUser.Username, newUser.Alias)
			Expect(err).ToNot(HaveOccurred())

			Expect(parsedConfig.Users).To(HaveLen(3))
			Expect(parsedConfig.Users[2]).To(Equal(newUser))
		})

		DescribeTable("validating errors get returned when",
			func(setup func() (User, Config, error)) {
				newUser, config, expectedError := setup()
				err := config.AddUser(newUser.ID, newUser.APIToken, newUser.Name, newUser.Username, newUser.Alias)
				Expect(err).To(MatchError(expectedError))
			},

			Entry("new user has the same alias as a previous user", func() (User, Config, error) {
				user := User{ID: 45, Username: "dventure", Alias: "dv"}
				newUser := User{ID: 46, Username: "dhventure", Alias: user.Alias}
				config := Config{Users: []User{user}}
				err := DuplicateAliasError{User: user}
				return newUser, config, err
			}),

			Entry("new user has the same username as a previous user's alias", func() (User, Config, error) {
				user := User{ID: 45, Username: "dventure", Alias: "dv"}
				newUser := User{ID: 46, Username: user.Alias, Alias: "dunmatter"}
				config := Config{Users: []User{user}}
				err := UsernameMatchesSavedAliasError{SavedUser: user, NewUser: newUser}
				return newUser, config, err
			}),

			Entry("new user has the same alias as a previous user's username", func() (User, Config, error) {
				user := User{ID: 45, Username: "dventure", Alias: "dv"}
				newUser := User{ID: 46, Username: "dhventure", Alias: user.Username}
				config := Config{Users: []User{user}}
				err := AliasMatchesSavedUsernameError{SavedUser: user, NewUser: newUser}
				return newUser, config, err
			}),
		)
	})

	Describe("SetCurrentUser", func() {
		Context("when user exists", func() {
			It("sets the current user with the current set time", func() {
				currentUser := parsedConfig.Users[0]
				err := parsedConfig.SetCurrentUser(currentUser.Username)
				Expect(err).ToNot(HaveOccurred())

				Expect(parsedConfig.CurrentUserID).To(Equal(currentUser.ID))
				Expect(parsedConfig.CurrentUserSetTime).To(BeTemporally("~", time.Now(), time.Second))
			})
		})

		Context("when user does not exist", func() {
			It("errors", func() {
				err := parsedConfig.SetCurrentUser("banana")
				Expect(err).To(Equal(ErrorUserDoesNotExist))
			})
		})

		Context("when passed an empty value", func() {
			It("sets current user and current set time to default values", func() {
				err := parsedConfig.SetCurrentUser("")
				Expect(err).ToNot(HaveOccurred())

				Expect(parsedConfig.CurrentUserID).To(BeZero())
				Expect(parsedConfig.CurrentUserSetTime).To(Equal(time.Time{}))
			})
		})
	})

	Describe("LoadConfig", func() {
		It("loads up the config", func() {
			config, err := LoadConfig(rawConfig)
			Expect(err).NotTo(HaveOccurred())
			Expect(config).To(Equal(parsedConfig))
		})
	})

	Describe("SaveConfig", func() {
		It("saves config based", func() {
			output, err := SaveConfig(parsedConfig)
			Expect(err).NotTo(HaveOccurred())
			Expect(output).To(MatchJSON(rawConfig))
		})
	})

})
