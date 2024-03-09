package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

type ApiError struct {
	Error string
}

/*
* Declared this type and the makeHttpHandleFunc function because we need to deal with the error
* the core func HandlerFunc doesn't support the error so we need to convert it
* to be able to use inside the mux
 */
type apiFunc func(http.ResponseWriter, *http.Request) error

func makeHTTPHandleFunc(f apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := f(w, r)
		if err != nil {
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

type APIServer struct {
	listenAddr string
}

func NewApiServer(listenAddr string) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
	}
}

func (s *APIServer) Start() {
	router := mux.NewRouter()
	router.HandleFunc("/account", makeHTTPHandleFunc(s.handleAccount))

	log.Println("Server running on port: ", s.listenAddr)

	http.ListenAndServe(s.listenAddr, router)
}

func (s *APIServer) handleAccount(w http.ResponseWriter, req *http.Request) error {

	switch req.Method {
	case "GET":
		return s.handleGetAccount(w, req)
	case "POST":
		s.handleCreateAccount(w, req)
	case "DELETE":
		s.handleDeleteAccount(w, req)
	}

	return fmt.Errorf("method not allowed %s", req.Method)
}

func (s *APIServer) handleGetAccount(w http.ResponseWriter, req *http.Request) error {
	return nil
}

func (s *APIServer) handleCreateAccount(w http.ResponseWriter, req *http.Request) error {

	return nil
}

func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, req *http.Request) error {

	return nil
}
