package config

import (
	"sort"
	"strings"

	"github.com/spf13/viper"
)

type User struct {
	Email    string
	ID       int64
	Name     string
	Token    string
	Username string
}

func (Config) GetUsers() ([]User, error) {
	users := []User{}
	for _, user := range viper.GetStringMap("users") {
		users = append(users, user.(User))
	}

	sort.Slice(users,
		func(i, j int) bool {
			return strings.Compare(users[i].Username, users[j].Username) == -1
		},
	)

	return users, nil
}
