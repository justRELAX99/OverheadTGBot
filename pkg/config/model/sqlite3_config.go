package model

type SqliteConfig struct {
	DBName       string `json:"dbName"`
	MaxOpenConns int    `json:"maxOpenConns"`
}
