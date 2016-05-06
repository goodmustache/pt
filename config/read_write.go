package config

import (
	"io/ioutil"
	"os"
	"path"
	"runtime"
)

func ReadConfig() (Config, error) {
	rawConfig, err := ioutil.ReadFile(configFile())
	if err != nil {
		return Config{}, err
	}

	return LoadConfig(rawConfig)
}

func WriteConfig(conf Config) error {
	rawConfig, err := SaveConfig(conf)
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
