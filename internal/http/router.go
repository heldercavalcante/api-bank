package http

import (
	"github.com/go-chi/chi/v5"
	userManagementRoutes "github.com/heldercavalcante/api-bank/internal/domain/user_management/routes"
)


func RegisterRoutes(router *chi.Mux) {
	userManagementRoutes.RegisterAccountRoutes(router)

}