package services

import (
	"tanam-backend/repository"
)

type AuthService struct {
	repository repository.AuthRepository
}

func InitAuthService() AuthService {
	return AuthService{
		repository: repository.InitAuthRepository(),
	}
}

func (service *AuthService) LoginService() (string, error) {
	service.repository.FindEmail("email")
	return "", nil
}
