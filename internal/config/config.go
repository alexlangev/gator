package config

import (
	"encoding/json"
	"errors"
	"os"
)

const configFile string = ".gatorconfig.json"

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (c *Config) SetUser(userName string) error {
	c.CurrentUserName = userName

	encoded, err := json.Marshal(c)
	if err != nil {
		return err
	}

	fileName, _ := getConfigFilePath()

	if err := os.WriteFile(fileName, encoded, 0755); err != nil {
		return err
	}

	return nil
}

func Read() (Config, error) {
	var config Config

	configFilePath, _ := getConfigFilePath()

	file, err := os.ReadFile(configFilePath)
	if err != nil {
		return config, errors.New("Error reading config file")
	}

	if err := json.Unmarshal(file, &config); err != nil {
		return config, errors.New("Error unmarshaling config file")
	}

	return config, nil
}

func getConfigFilePath() (string, error) {
	userHome, err := os.UserHomeDir()
	if err != nil {
		return "", errors.New("Can't find user home!")
	}

	return userHome + "/" + configFile, nil
}
