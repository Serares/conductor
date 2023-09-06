package server

import (
	"net/http"

	"github.com/Serares/conductor/repository"
)

type HttpServer struct {
	server *http.Server
	repo   *repository.Repository
}

type RegisterResponse struct {
}

type RegisterRequest struct {
	Ip              string `json:"ip"`
	TransactionHash string `json:"transaction_hash"`
}
