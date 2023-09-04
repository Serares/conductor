package server

import (
	"net/http"

	"github.com/Serares/conductor/repository"
)

type HttpServer struct {
	server http.Server
	repo   repository.Repository
}
