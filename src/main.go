package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"

	"github.com/gnemes/golang-ws-skeleton/Domain/Services/Logger"
	"github.com/gnemes/golang-ws-skeleton/Infrastructure/Services/Container"
)

func main() {
	logger := container.GetInstance().Get("Logger").(logger.Logger)
	logger.Infof("----- Starting Web Server... -----")
	defer logger.Infof("----- Ending Web Server -----")
	defer container.DeleteInstance()
	//defer database.Close()
	
	s := &Server{
		Router: mux.NewRouter().StrictSlash(true),
	}

	log.Fatal(http.ListenAndServe(":8080", s.Routes()))
}