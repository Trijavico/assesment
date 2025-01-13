package controller

import (
	"encoding/json"
	"net/http"

	"github.com/Trijavico/ensolvers-assesment/internal/model/dto"
	"github.com/Trijavico/ensolvers-assesment/internal/service"
)

type AuthController struct {
	authService *service.AuthService
}

func NewAuthController(authService *service.AuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

func (auth *AuthController) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var createUser dto.CreateUser

	err := json.NewDecoder(r.Body).Decode(&createUser)
	if err != nil {
		http.Error(w, "Unable to create account", http.StatusInternalServerError)
		return
	}

	token, err := auth.authService.Signup(w, r, createUser)
	if err != nil {
		http.Error(w, "Unable to create account", http.StatusInternalServerError)
		return
	}

	authenticated := struct {
		Token string `json:"token"`
	}{
		Token: token,
	}

	err = RespondJSON(w, r, http.StatusCreated, authenticated)
	if err != nil {
		http.Error(w, "Unable to create account", http.StatusInternalServerError)
	}
}

func (auth *AuthController) LogInAccount(w http.ResponseWriter, r *http.Request) {
	var loginUser dto.LoginUser

	err := json.NewDecoder(r.Body).Decode(&loginUser)
	if err != nil {
		http.Error(w, "Unable to log into account", http.StatusInternalServerError)
		return
	}

	token, err := auth.authService.Login(w, r, loginUser)
	if err != nil {
		http.Error(w, "Unable to log into account", http.StatusInternalServerError)
		return
	}

	authenticated := struct {
		Token string `json:"token"`
	}{
		Token: token,
	}

	err = RespondJSON(w, r, http.StatusCreated, authenticated)
	if err != nil {
		http.Error(w, "Unable to log into account", http.StatusInternalServerError)
	}
}
