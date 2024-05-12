package repository

import (
	"tanam-backend/database"
	"tanam-backend/domains/model"
	"tanam-backend/domains/web/auth"
	"tanam-backend/entities"

	"github.com/google/uuid"
)

type BiodataRepository struct {
}

func InitBiodataRepository() BiodataRepository {
	return BiodataRepository{}
}

func (repo *BiodataRepository) CreateBiodata(request auth.RegisterRequest, authId uuid.UUID) error {
	biodata := model.Biodata{
		BiodataID:   uuid.New(),
		AuthID:      authId,
		Fullname:    request.Fullname,
		PhoneNumber: request.PhoneNumber,
	}

	if err := database.DB.Create(&biodata).Error; err != nil {
		return err
	}

	return nil
}

func (repo *BiodataRepository) FIndBiodataByAuthId(authId string) (entities.Biodata, error) {
	var biodata entities.Biodata
	if err := database.DB.Where("auth_id = ?", &authId).First(&biodata).Error; err != nil {
		return biodata, err
	}
	return biodata, nil
}
