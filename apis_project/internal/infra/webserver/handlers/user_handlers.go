package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/bruno-holanda15/go_expert_fc/APIs_project/internal/dto"
	"github.com/bruno-holanda15/go_expert_fc/APIs_project/internal/entity"
	"github.com/bruno-holanda15/go_expert_fc/APIs_project/internal/infra/database"
	"github.com/go-chi/jwtauth"
)

type UserHandler struct {
	UserDB database.UserDBInterface
}

func NewUserHandler(db database.UserDBInterface) *UserHandler {
	return &UserHandler{
		UserDB: db,
	}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var userDTO dto.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&userDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := entity.NewUser(userDTO.Name, userDTO.Email, userDTO.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.UserDB.Create(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *UserHandler) GetJWT(w http.ResponseWriter, r *http.Request) {
	var infoInput dto.GetJWTInput
	err := json.NewDecoder(r.Body).Decode(&infoInput)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := h.UserDB.FindByEmail(infoInput.Email)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if !user.ValidatePassword(infoInput.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		return	
	}

	jwt := r.Context().Value("jwt").(jwtauth.JWTAuth)
	jwtExpiresIn := r.Context().Value("JwtExperesIn").(int)
	_, tokenString, _ := jwt.Encode(map[string]interface{}{
		"sub": user.ID.String(),
		"exp": time.Now().Add(time.Second * time.Duration(jwtExpiresIn)).Unix(),
	})

	accessToken := dto.GetJWTOutput{AccessToken: tokenString}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)
}