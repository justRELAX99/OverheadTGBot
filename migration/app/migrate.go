package app

import (
	"OverheadTGBot/pkg/config"
	"OverheadTGBot/pkg/errors"
	"OverheadTGBot/pkg/logger"
	"flag"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pressly/goose/v3"
)

const (
	driverName = "sqlite3"
)

func Run(configSettings config.Config) {
	logger := logger.NewZapLogger(configSettings.Logger)

	dbConfig := configSettings.Sqlite
	command := flag.String("c", "status", "command")
	args := flag.String("args", "", "args")
	dir := flag.String("dir", "./migration/migrations", "migration dir")
	flag.Parse()

	dsn := fmt.Sprintf("%v/%v.db", dbConfig.Path, dbConfig.DBName)

	db, err := goose.OpenDBWithDriver(driverName, dsn)
	if err != nil {
		panic(errors.Wrap(err, "goose: failed to open DB"))
	}
	defer func() {
		if err := db.Close(); err != nil {
			panic(errors.Wrap(err, "goose: failed to close DB"))
		}
	}()
	if err := goose.Run(*command, db, *dir, *args); err != nil {
		panic(errors.Wrap(err, "goose: failed run migrations"))
	}
	logger.Info("Migrations end")
}
