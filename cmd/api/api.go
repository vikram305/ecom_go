package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	user "github.com/vikram305/ecom/services"
)

type APIServer struct {
	address string
	db      *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		address: addr,
		db:      db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userHandler := user.NewHandler()
	userHandler.RegisterRoutes(subrouter)

	log.Println("Listening on ", s.address)
	return http.ListenAndServe(s.address,router)
}
