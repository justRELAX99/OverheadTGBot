package config

import (
	"OverheadTGBot/pkg"
	"OverheadTGBot/pkg/config/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Config struct {
	TelegramBot model.TelegramBotConfig `json:"telegramBot"`
	Logger      model.LoggerConfig      `json:"logger"`
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
