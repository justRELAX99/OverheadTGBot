package main

import (
	"OverheadTGBot/migration/app"
	"OverheadTGBot/pkg/config"
)

const configsDirectory = "configs"

func main() {
	app.Run(config.LoadConfigSettings(configsDirectory))
}
