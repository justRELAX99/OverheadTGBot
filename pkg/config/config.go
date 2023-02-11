package config

import (
	"OverheadTGBot/pkg"
	"OverheadTGBot/pkg/config/entity"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Config struct {
	TelegramBot entity.TelegramBotConfig `json:"telegramBot"`
	Logger      entity.LoggerConfig      `json:"logger"`
	Sqlite      entity.SqliteConfig      `json:"sqlite"`
}

func LoadConfigSettings(configsDirectory string) (config Config) {
	environment := pkg.GetEnvironment()
	path := fmt.Sprintf("%v/%v.json", configsDirectory, environment)
	byteValue, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		log.Fatal(err)
	}
	return config
}
