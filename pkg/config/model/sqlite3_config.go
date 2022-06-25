package model

type SqliteConfig struct {
	Path         string `json:"path"`
	DBName       string `json:"dbName"`
	MaxOpenConns int    `json:"maxOpenConns"`
}
