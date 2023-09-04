package server

import "github.com/Serares/conductor/repository"

func getRepo() (repository.Repository, error) {
	repo, err := repository.NewPsqlRepo()
	if err != nil {
		return nil, err
	}
	return repo, nil
}
