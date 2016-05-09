package actions

import (
	"io/ioutil"
	"os"
	"path"
	"runtime"

	"github.com/goodmustache/pt/config"
)

// ReadConfig reads the config in $HOME/.config/pt/config.json.
func ReadConfig() (config.Config, error) {
	rawConfig, err := ioutil.ReadFile(configFile())
	if err != nil {
		return config.Config{}, err
	}

	return config.LoadConfig(rawConfig)
}

// WriteConfig writes the config in $HOME/.config/pt/config.json, if no such
// file/directory exists, it will create them.
func WriteConfig(conf config.Config) error {
	rawConfig, err := config.SaveConfig(conf)
	if err != nil {
		return err
	}

	if _, err := os.Stat(configFile()); os.IsNotExist(err) {
		err = os.MkdirAll(configDir(), 0777)
		if err != nil {
			return err
		}
	}

	file, err := os.Create(configFile())
	if err != nil {
		return err
	}
	_, err = file.Write(rawConfig)

	return err
}

func configFile() string {
	return path.Join(configDir(), "config.json")
}

func configDir() string {
	return path.Join(userHomeDir(), ".config", "pt")
}

func userHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}

	return os.Getenv("HOME")
}
