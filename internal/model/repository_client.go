package model

import (
	"github.com/gocraft/dbr"
)

type RepositoryClient interface {
	GetSession() dbr.SessionRunner
}
