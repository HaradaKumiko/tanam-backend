package services

import (
	"errors"
	"tanam-backend/domains/web/auth"
	"tanam-backend/entities"
	"tanam-backend/helpers"
	"tanam-backend/repository"
)

type AuthService struct {
	authRepository    repository.AuthRepository
	biodataRepository repository.BiodataRepository
}

func InitAuthService() AuthService {
	return AuthService{
		authRepository:    repository.InitAuthRepository(),
		biodataRepository: repository.InitBiodataRepository(),
	}
}

func (service *AuthService) LoginService(loginRequest auth.LoginRequest) (string, error) {
	auth, err := service.authRepository.FindEmail(loginRequest.Email)
	if err != nil {
		return "", errors.New("email tidak ditemukan")
	}

	if !helpers.CheckPasswordHash(loginRequest.Password, auth.Password) {
		return "", errors.New("password salah")
	}

	token, err := helpers.GenerateTokenJWT(auth.AuthID, auth.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (service *AuthService) RegisterService(registerRequest auth.RegisterRequest) (entities.Auth, error) {

	emailExists, err := service.authRepository.EmailExists(registerRequest.Email)
	if err != nil {
		return entities.Auth{}, err
	}
	if emailExists {
		return entities.Auth{}, errors.New("email already exists")
	}

	hashedPassword := helpers.HashPassword(registerRequest.Password)
	registerRequest.Password = hashedPassword

	createdAuth, err := service.authRepository.CreateAuth(registerRequest)
	if err != nil {
		return entities.Auth{}, err
	}

	err = service.biodataRepository.CreateBiodata(registerRequest, createdAuth.AuthID)
	if err != nil {
		return entities.Auth{}, err

	}

	return createdAuth, nil
}

func (service *AuthService) ProfileService(authId string) (entities.Auth, error) {
	auth, err := service.authRepository.FindAuthById(authId)
	if err != nil {
		return auth, err
	}

	return auth, nil
}
