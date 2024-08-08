package http

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type APIServer struct {
	Port string
}

func NewAPIServer(port string) *APIServer {
	return &APIServer{
		Port: port,
	}
}

func (server *APIServer) Run() {
	router := chi.NewRouter()
	RegisterRoutes(router)

	fmt.Println("Running server on port: ", server.Port)
	http.ListenAndServe(fmt.Sprintf(":%s", server.Port), router)
}