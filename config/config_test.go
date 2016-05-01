package config_test

import (
	"time"

	. "github.com/goodmustache/pt/config"

	. "github.com/onsi/ginkgo"
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
				"aliases": ["dv"]
			},
			{
				"id": 46,
				"username": "hventure",
				"name": "Henry Allen 'Hank' Venture",
				"api_token": "22222222222222222222222222222222",
				"aliases": ["hv", "hav"]
			}
		]
	}`)

		parsedConfig = Config{
			CurrentUserID:      42,
			CurrentUserSetTime: time.Date(2016, 4, 26, 7, 37, 47, 0, time.UTC),
			Users: []User{
				{
					ID:       45,
					APIToken: "11111111111111111111111111111111",
					Name:     "Dean Venture",
					Username: "dventure",
					Aliases:  []string{"dv"},
				},
				{
					ID:       46,
					APIToken: "22222222222222222222222222222222",
					Name:     "Henry Allen 'Hank' Venture",
					Username: "hventure",
					Aliases:  []string{"hv", "hav"},
				},
			},
		}
	})

	Describe("AddUser", func() {
		var user User
		BeforeEach(func() {
			user = User{
				ID:       47,
				APIToken: "ffffffffffffffffffffffffffffffff",
				Name:     "Dermott Fictel",
				Username: "TheWolf",
				Aliases:  []string{"tw", "fs"},
			}
		})

		It("adds the user", func() {
			err := parsedConfig.AddUser(user.ID, user.APIToken, user.Name, user.Username, user.Aliases)
			Expect(err).ToNot(HaveOccurred())

			Expect(parsedConfig.Users).To(HaveLen(3))
			Expect(parsedConfig.Users[2]).To(Equal(user))
		})

		It("overrides user if id already exists", func() {
			user.ID = 46

			err := parsedConfig.AddUser(user.ID, user.APIToken, user.Name, user.Username, user.Aliases)
			Expect(err).ToNot(HaveOccurred())

			Expect(parsedConfig.Users).To(HaveLen(2))
			Expect(parsedConfig.Users[1]).To(Equal(user))
		})

		It("errors when alias already exists", func() {
			originalUser := parsedConfig.Users[0]
			user.Aliases = append(user.Aliases, originalUser.Aliases[0])

			err := parsedConfig.AddUser(user.ID, user.APIToken, user.Name, user.Username, user.Aliases)
			Expect(err).To(MatchError(DuplicateAliasError{User: originalUser}))
		})
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
				Expect(err).To(Equal(UserDoesNotExistError))
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
