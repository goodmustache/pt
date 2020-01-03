package config

import (
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/goodmustache/pt/actor"
	"github.com/spf13/viper"
)

type Config struct{}

func SetDefaultConfig() {
	viper.SetDefault("users", map[string]actor.User{})
}

func ReadConfig() (*Config, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	viper.SetConfigName("config")
	viper.AddConfigPath(filepath.Join(homeDir, ".config", "pt"))
	return new(Config), viper.ReadInConfig()
}

func WriteConfig() error {
	return viper.WriteConfig()
}

func (Config) GetUsers() ([]actor.User, error) {
	users := []actor.User{}
	for _, user := range viper.GetStringMap("users") {
		users = append(users, user.(actor.User))
	}

	sort.Slice(users,
		func(i, j int) bool {
			return strings.Compare(users[i].Username, users[j].Username) == -1
		},
	)

	return users, nil
}
