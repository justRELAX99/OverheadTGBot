package main

import (
	"OverheadTGBot/internal/app"
	"OverheadTGBot/pkg/config"
)

const configsDirectory = "configs"

func main() {
	app.Run(config.LoadConfigSettings(configsDirectory))
}
