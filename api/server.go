package api

import (
	"encoding/json"
	"apiServer/storage"
	"net/http"
)

type Server struct{
	listenAddr 	string
	store 		storage.Storage
}

// constructor
func NewServer(listenAddr string, store storage.Storage) *Server{
	return &Server{
		listenAddr: listenAddr,
		store: store,
	}
}


func (s *Server) Start() error {
	http.HandleFunc("/user", s.handleGetUserByID)
	http.HandleFunc("/user/createNewUser", s.handleCreateNewUser)
	return http.ListenAndServe(s.listenAddr, nil)
}

// handler for /user api calls
func (s *Server) handleGetUserByID(w http.ResponseWriter, r *http.Request){
	user := s.store.Get(10)

	json.NewEncoder(w).Encode(user)
}

//handler for user create calls
func (s *Server) handleCreateNewUser(w http.ResponseWriter, r *http.Request){
	user := s.store.Get(10)

	json.NewEncoder(w).Encode(user)
}