package service

import (
	"context"
	"net/http"

	"github.com/Trijavico/ensolvers-assesment/internal/model"
	"github.com/Trijavico/ensolvers-assesment/internal/model/dto"
	"github.com/Trijavico/ensolvers-assesment/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

var Session_cookie = "sess"

type AuthService struct {
	userRepo   repository.UserRepository
	jwtService *JwtService
}

func NewAuthService(userRepo repository.UserRepository, jwtService *JwtService) *AuthService {
	return &AuthService{
		userRepo:   userRepo,
		jwtService: jwtService,
	}
}

func (service *AuthService) Signup(w http.ResponseWriter, r *http.Request, createUser dto.CreateUser) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(createUser.Password), 12)
	if err != nil {
		return "", err
	}

	user := &model.User{
		Email:    createUser.Email,
		Password: hashedPassword,
	}

	err = service.userRepo.Save(context.Background(), user)
	if err != nil {
		return "", err
	}

	claims := map[string]interface{}{
		"id": user.ID,
	}
	token, err := service.jwtService.CreateToken(claims)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (service *AuthService) Login(w http.ResponseWriter, r *http.Request, loginUser dto.LoginUser) (string, error) {
	result, err := service.userRepo.FindByEmail(context.Background(), loginUser.Email)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword(result.Password, []byte(loginUser.Password))
	if err != nil {
		return "", err
	}

	claims := map[string]interface{}{
		"id": result.ID,
	}

	token, err := service.jwtService.CreateToken(claims)
	if err != nil {
		return "", err
	}

	return token, nil
}
