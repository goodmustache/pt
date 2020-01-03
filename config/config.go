package config

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/spf13/viper"
)

type Config struct{}

func ConfigDir() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("Error determining user's home directory: %w", err)
	}

	return filepath.Join(homeDir, ".config", "pt"), nil
}

func InitConfig(configDir string) error {
	_, err := os.Stat(configDir)

	if os.IsNotExist(err) {
		err = os.MkdirAll(configDir, 0755)
		if err != nil {
			return err
		}

		cfgPath := filepath.Join(configDir, "config.json")
		f, err := os.Create(cfgPath)
		if err != nil {
			return err
		}
		f.Close()
	}

	return err
}

func ReadConfig() (*Config, error) {
	configDir, err := ConfigDir()
	if err != nil {
		return nil, err
	}

	viper.SetConfigName("config")
	viper.AddConfigPath(configDir)
	err = viper.ReadInConfig()
	if _, ok := err.(viper.ConfigFileNotFoundError); !ok && err != nil {
		return nil, err
	}

	return new(Config), nil
}

func SetDefaultConfig() {
	viper.SetDefault("users", map[string]User{})
}

func WriteConfig() error {
	configDir, err := ConfigDir()
	if err != nil {
		return err
	}

	err = InitConfig(configDir)
	if err != nil {
		return err
	}

	return viper.WriteConfig()
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
