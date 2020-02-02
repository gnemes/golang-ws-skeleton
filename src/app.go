package main

import (
	"github.com/gorilla/mux"
)

type Server struct {
	Router *mux.Router
}

func (s *Server) Routes() *mux.Router {
	// s.Router.Use(middleware.ApplicationJsonMiddleware)
	//s.Router.Use(middleware.RequestIDMiddleware)

	// Health check
	//s.Router.HandleFunc("/healthcheck", healthactions.Check).Methods("GET", "POST")

	return s.Router
}