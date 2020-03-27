package config

import (
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/spf13/viper"
)

const UsersKey = "users"

type ConfigUsers map[string]User

type User struct {
	APIToken string
	Email    string
	ID       uint64
	Name     string
	Username string
}

func (Config) GetAPITokenForUser(id uint64) (string, error) {
	var users ConfigUsers
	err := viper.UnmarshalKey("users", &users)
	if err != nil {
		return "", err
	}

	if len(users) == 0 {
		return "", errors.New("no configured users")
	}

	if user, ok := users[fmt.Sprint(id)]; ok {
		return user.APIToken, nil
	}

	return "", errors.New("provided user id not found")
}

func (Config) GetUsers() ([]User, error) {
	var raw ConfigUsers
	err := viper.UnmarshalKey("users", &raw)
	if err != nil {
		return nil, err
	}

	users := []User{}
	for _, user := range raw {
		users = append(users, user)
	}

	sort.Slice(users,
		func(i, j int) bool {
			return strings.Compare(users[i].Username, users[j].Username) == -1
		},
	)

	return users, nil
}

func (Config) AddUser(user User) error {
	var raw ConfigUsers
	err := viper.UnmarshalKey("users", &raw)
	if err != nil {
		return err
	}

	raw[fmt.Sprint(user.ID)] = user

	viper.Set("users", raw)
	return nil
}
