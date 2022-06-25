package sqlite

import (
	"OverheadTGBot/internal/model"
	config "OverheadTGBot/pkg/config/model"
	"fmt"
	"github.com/gocraft/dbr"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"time"
)

const (
	driverName = "sqlite3"
)

type sqliteClient struct {
	config     config.SqliteConfig
	connection *dbr.Connection
	logger     model.Logger
}

func NewClient(config config.SqliteConfig, logger model.Logger) model.RepositoryClient {
	client := sqliteClient{
		config: config,
		logger: logger,
	}
	ok := client.setConnection()
	if !ok {
		log.Fatal("Cant set sqlite connect")
	}
	return &client
}

func (c sqliteClient) GetSession() dbr.SessionRunner {
	return c.getConnection().NewSession(&dbr.NullEventReceiver{})
}

func (c *sqliteClient) setConnection() (ok bool) {
	connectionString := fmt.Sprintf("%v.db", c.config.DBName)

	conn, err := dbr.Open(driverName, connectionString, nil)
	if err != nil {
		c.logger.Error("An error occurred while initializing the connection to postgres:", err.Error())
		return false
	}
	conn.SetMaxOpenConns(c.config.MaxOpenConns)
	conn.SetMaxIdleConns(2)
	conn.SetConnMaxIdleTime(10 * time.Second) //обрубаем конекшн, если 10 секунд ничего не делал

	c.connection = conn
	if !c.checkConnection() {
		c.logger.Error("Postgres connection not ready")
		return false
	}
	return true
}

func (c sqliteClient) getConnection() *dbr.Connection {
	if !c.checkConnection() {
		c.reconnect()
	}
	return c.connection
}

func (c sqliteClient) checkConnection() (ok bool) {
	if c.connection == nil {
		return false
	}
	err := c.connection.Ping()
	if err != nil {
		c.logger.Error("Cant ping connection,", err.Error())
		return false
	}

	return true
}

func (c *sqliteClient) reconnect() {
	if c.connection != nil {
		err := c.connection.Close()
		if err != nil {
			c.logger.Debug("Connection close error,", err.Error())
		}
	}

	c.logger.Debug("start reconnect")
	for {
		time.Sleep(time.Second / 10)
		ok := c.setConnection()
		if !ok {
			continue
		}
		return
	}
}
