package repository

import (
	"database/sql"

	"github.com/go-pg/pg/v10"
)

type dbRepo struct {
	db *sql.DB
}

type ConnectionOptions struct {
	Username string
	Password string
	Hostname string
	Database string
}

func NewPsqlRepo(options ConnectionOptions) (*pg.DB, error) {
	db := pg.Connect(&pg.Options{
		Password: options.Password,
		User:     options.Username,
		Addr:     options.Hostname,
		Database: options.Database,
	})

}
