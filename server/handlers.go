package server

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func v1Handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/register" {
			if r.Method == http.MethodPost {
				register(w, r)
			}
		}

		if r.URL.Path == "/file" {
			if r.Method == http.MethodPost {
				fmt.Println("Send the file")
			}
		}
	}
}

func register(w http.ResponseWriter, r *http.Request) {
	// timestamp := r.Header.Get("x-relayr-timestamp")
	// secrets := r.Header.Get("x-relayr-secrets")

	// decoded, err := base64.StdEncoding.DecodeString(secrets)
	response, err := http.Post("http://localhost:8080/v1/file", "application/json", nil)
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
	w.Header().Set("x-relayr-timestamp", fmt.Sprintf("%v", time.Now()))
	w.WriteHeader(status)
	w.Write([]byte(content))
}

func replyError(w http.ResponseWriter, r *http.Request, status int, message string) {
	log.Printf("%s %s: Error: %d %s", r.URL, r.Method, status, message)
	http.Error(w, http.StatusText(status), status)
}
