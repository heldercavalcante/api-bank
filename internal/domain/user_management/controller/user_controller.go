package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/heldercavalcante/api-bank/internal/domain/user_management/entity"
	"github.com/heldercavalcante/api-bank/internal/domain/user_management/service"
	"github.com/heldercavalcante/api-bank/internal/tools"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	req := new(entity.UserRequest)
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("Error trying to make the json decode: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	validation, err := service.Validate(req);

	if err != nil {
		log.Printf("Validation failed: %v", err)
		response := map[string]interface{}{
			"error": map[string]interface{}{
				"type": "INTERNAL_ERROR",
				"message": http.StatusText(http.StatusInternalServerError),
				"description": "the user couldn't be created because of internal problems!",
			},
		}
		tools.WriteResponseJSON(w, http.StatusInternalServerError, response)
		return
	}
	if !validation.Valid {
		response := map[string]interface{}{
			"error": map[string]interface{}{
				"type": "VALIDATION_ERROR",
				"message": validation.Message,
				"description": "the user couldn't be created because of validation problems!",
			},
		}
		
		tools.WriteResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	userId, err := service.CreateUser(*req)
	if err != nil {
		log.Printf("Error trying to create a user: %v", err)
		response := map[string]interface{}{
			"error": map[string]interface{}{
				"type": "INTERNAL_ERROR",
				"message": http.StatusText(http.StatusInternalServerError),
				"description": "the user couldn't be created because of internal problems!",
			},
		}
		tools.WriteResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	user, err := service.GetUserCompleteDataById(int32(userId))
	if err != nil {
		log.Printf("Error trying to get a user: %v", err)
		response := map[string]interface{}{
			"error": map[string]interface{}{
				"type": "BAD_REQUEST",
				"message": http.StatusText(http.StatusBadRequest),
				"description": "Couldn't find the user with the id provided",
			},
		}
		tools.WriteResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	tools.WriteResponseJSON(w, http.StatusOK, user)
}


func GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := service.GetUsersCompleteData()
	if err != nil {
		log.Printf("Error trying to get users: %v", err)
		response := map[string]interface{}{
			"error": map[string]interface{}{
				"type": "INTERNAL_ERROR",
				"message": http.StatusText(http.StatusInternalServerError),
				"description": "Couldn't find any user",
			},
		}
		tools.WriteResponseJSON(w, http.StatusInternalServerError, response)
		return
	}
	tools.WriteResponseJSON(w, http.StatusOK, users)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	userId, err := strconv.Atoi(idParam) 
	if err != nil {
		log.Printf("Invalid user ID: %v", err)
		response := map[string]interface{}{
			"error": map[string]interface{}{
				"type": "BAD_REQUEST",
				"message": http.StatusText(http.StatusBadRequest),
				"description": "Invalid user ID",
			},
		}
		tools.WriteResponseJSON(w, http.StatusBadRequest, response)
		return
    }
	user, err := service.GetUserCompleteDataById(int32(userId))

	if err != nil {
		log.Printf("Error trying to get a user: %v", err)
		response := map[string]interface{}{
			"error": map[string]interface{}{
				"type": "BAD_REQUEST",
				"message": http.StatusText(http.StatusBadRequest),
				"description": "Couldn't find the user with the id provided",
			},
		}
		tools.WriteResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	tools.WriteResponseJSON(w, http.StatusOK, user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	userId, err := strconv.Atoi(idParam) 
	if err != nil {
		log.Printf("Invalid user ID: %v", err)
		response := map[string]interface{}{
			"error": map[string]interface{}{
				"type": "BAD_REQUEST",
				"message": http.StatusText(http.StatusBadRequest),
				"description": "Invalid user ID",
			},
		}
		tools.WriteResponseJSON(w, http.StatusBadRequest, response)
		return
    }
	err = service.DeleteUser(int32(userId))

	if err != nil {
		log.Printf("Error trying to delete a user: %v", err)
		response := map[string]interface{}{
			"error": map[string]interface{}{
				"type": "BAD_REQUEST",
				"message": http.StatusText(http.StatusBadRequest),
				"description": "Couldn't find the user with the id provided",
			},
		}
		tools.WriteResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	tools.WriteResponseJSON(w, http.StatusOK, "User successfully deleted")
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "id")
	userId, err := strconv.Atoi(idParam) 
	if err != nil {
		log.Printf("Invalid user ID: %v", err)
		response := map[string]interface{}{
			"error": map[string]interface{}{
				"type": "BAD_REQUEST",
				"message": http.StatusText(http.StatusBadRequest),
				"description": "Invalid user ID",
			},
		}
		tools.WriteResponseJSON(w, http.StatusBadRequest, response)
		return
    }
	req := new(entity.UserRequest)
	req.UserID = int64(userId)
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("Error trying to make the json decode: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	validation, err := service.Validate(req);

	if err != nil {
		log.Printf("Validation failed: %v", err)
		response := map[string]interface{}{
			"error": map[string]interface{}{
				"type": "INTERNAL_ERROR",
				"message": http.StatusText(http.StatusInternalServerError),
				"description": "the user couldn't be Updated because of internal problems!",
			},
		}
		tools.WriteResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	if !validation.Valid {
		response := map[string]interface{}{
			"error": map[string]interface{}{
				"type": "VALIDATION_ERROR",
				"message": validation.Message,
				"description": "the user couldn't be Updated because of validation problems!",
			},
		}
		
		tools.WriteResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	err = service.UpdateUser(*req)
	if err != nil {
		log.Printf("Error trying to Update a user: %v", err)
		response := map[string]interface{}{
			"error": map[string]interface{}{
				"type": "BAD_REQUEST",
				"message": http.StatusText(http.StatusBadRequest),
				"description": "Couldn't find the user with the id provided",
			},
		}
		tools.WriteResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	user, err := service.GetUserCompleteDataById(int32(userId))
	if err != nil {
		log.Printf("Error trying to get a user: %v", err)
		response := map[string]interface{}{
			"error": map[string]interface{}{
				"type": "BAD_REQUEST",
				"message": http.StatusText(http.StatusBadRequest),
				"description": "Couldn't find the user with the id provided",
			},
		}
		tools.WriteResponseJSON(w, http.StatusBadRequest, response)
		return
	}

	tools.WriteResponseJSON(w, http.StatusOK, user)
}