package server

import (
	"os"

	"github.com/Serares/conductor/repository"
)

func getRepo() (repository.Repository, error) {
	username := os.Getenv("PSQL_USERNAME")
	password := os.Getenv("PSQL_PASSWORD")
	db := os.Getenv("PSQL_DB")
	host := os.Getenv("PSQL_HOST")
	repo, err := repository.NewPsqlRepo(repository.ConnectionOptions{Username: username, Password: password, Database: db, Hostname: host})
	if err != nil {
		return nil, err
	}
	return repo, nil
}
