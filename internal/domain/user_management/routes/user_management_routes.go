package routes


import (
	"github.com/go-chi/chi/v5"
	userManagementController "github.com/heldercavalcante/api-bank/internal/domain/user_management/controller"
)

func RegisterAccountRoutes(router *chi.Mux) {
	router.Post("/user", userManagementController.CreateUser)
	router.Get("/user/{id}", userManagementController.GetUser)
	router.Get("/user", userManagementController.GetUsers)
	router.Delete("/user/{id}", userManagementController.DeleteUser)
	router.Put("/user/{id}", userManagementController.UpdateUser)
}