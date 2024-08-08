package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/heldercavalcante/api-bank/internal/domain/user_management/entity"
	"github.com/heldercavalcante/api-bank/internal/domain/user_management/service"
	"github.com/heldercavalcante/api-bank/internal/tools"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	req := new(entity.CreateUserRequest)
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("Erro ao fazer decode do json: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	userId, err := service.CreateUser(*req)
	if err != nil {
		log.Printf("Erro na tentativa de criar o usuario: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	tools.WriteResponseJSON(w, http.StatusOK, map[string]interface{}{"userId": userId})
}

//teste
func GetUser(w http.ResponseWriter, r *http.Request) {
	tools.WriteResponseJSON(w, http.StatusOK, "retorno coerente")
}