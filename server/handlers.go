package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Serares/conductor/repository"
	"github.com/google/uuid"
)

func v1Handler(repo repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "register" {
			if r.Method == http.MethodPost {
				register(w, r, repo)
				return
			}
		}

		if r.URL.Path == "find" {
			if r.Method == http.MethodPost {

			}
		}

		replyError(w, r, http.StatusNotFound, fmt.Sprintf("Handler not found: %s, %s", r.URL.Path, r.Method))
	}
}

func register(w http.ResponseWriter, r *http.Request, repo repository.Repository) {
	registerRequest := &RegisterRequest{}

	if err := json.NewDecoder(r.Body).Decode(registerRequest); err != nil {
		message := fmt.Sprintf("Invalid JSON: %s", err)
		replyError(w, r, http.StatusBadRequest, message)
		return
	}
	uniqueId := uuid.New()
	peer := repository.PeerInfo{Id: uniqueId.String(), Ip: registerRequest.Ip, TransactionHash: registerRequest.TransactionHash}
	_, err := repo.Store(peer)
	if err != nil {
		replyError(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	replyTextContent(w, r, http.StatusAccepted, fmt.Sprintf("registered success: %s", r.RemoteAddr))
}

func initialiseHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			replyError(w, r, http.StatusNotFound, "Not found")
			return
		}
		content := "Route found"
		replyTextContent(w, r, http.StatusOK, content)
	}
}

func replyTextContent(w http.ResponseWriter, r *http.Request, status int, content string) {
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("x-conductor-timestamp", fmt.Sprintf("%v", time.Now()))
	w.WriteHeader(status)
	w.Write([]byte(content))
}

func replyError(w http.ResponseWriter, r *http.Request, status int, message string) {
	log.Printf("%s %s: Error: %d %s", r.URL, r.Method, status, message)
	http.Error(w, http.StatusText(status), status)
}
