package server

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/Serares/conductor/repository"
)

func newMux(repo repository.Repository) http.Handler {
	m := http.NewServeMux()

	v1 := v1Handler(repo)

	m.Handle("/v1/", http.StripPrefix("/v1/", v1))
	return m
}

func newHttp() (*HttpServer, error) {
	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		serverPort = "8080"
	}
	fmt.Println("Listening on port:", serverPort)
	repo, err := getRepo()
	if err != nil {
		return nil, err
	}
	server := &http.Server{
		Addr:         fmt.Sprintf("localhost:%s", serverPort),
		Handler:      newMux(repo),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	return &HttpServer{
		server: server,
		repo:   &repo,
	}, nil

}

func InitHttpServer() {
	s, err := newHttp()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if err := s.server.ListenAndServe(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
