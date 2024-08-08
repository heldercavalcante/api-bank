package routes


import (
	"github.com/go-chi/chi/v5"
	userManagementController "github.com/heldercavalcante/api-bank/internal/domain/user_management/controller"
)

func RegisterAccountRoutes(router *chi.Mux) {
	router.Post("/user", userManagementController.CreateUser)
	router.Get("/user", userManagementController.GetUser)
}