package repository

import (
	"tanam-backend/database"
	"tanam-backend/domains/web/auth"
	"tanam-backend/entities"

	"github.com/google/uuid"
)

type AuthRepository struct {
}

func InitAuthRepository() AuthRepository {
	return AuthRepository{}
}

func (repo *AuthRepository) FindEmail(email string) (entities.Auth, error) {
	var auth entities.Auth
	if err := database.DB.Where("email = ?", &email).First(&auth).Error; err != nil {
		return auth, err
	}
	return auth, nil
}

func (repo *AuthRepository) EmailExists(email string) (bool, error) {
    var count int64
    if err := database.DB.Model(&entities.Auth{}).Where("email = ?", email).Count(&count).Error; err != nil {
        return false, err
    }
    return count > 0, nil
}

func (repo *AuthRepository) FindAuthById(authId string) (entities.Auth, error) {
	var auth entities.Auth
	if err := database.DB.Where("auth_id = ?", &authId).First(&auth).Error; err != nil {
		return auth, err
	}
	return auth, nil
}

func (repo *AuthRepository) CreateAuth(request auth.RegisterRequest) (entities.Auth, error) {
	auth := entities.Auth{
		AuthID:   uuid.New(),
		Email:    request.Email,
		Password: request.Password,
		Profile:  "https://avatars.githubusercontent.com/u/42530587",
	}

	if err := database.DB.Create(&auth).Error; err != nil {
		return entities.Auth{}, err
	}

	return auth, nil
}
