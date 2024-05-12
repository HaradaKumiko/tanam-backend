package repository

import (
	"errors"
	"tanam-backend/database"
	"tanam-backend/domains/model"
	"tanam-backend/domains/web/plant"
	"tanam-backend/entities"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PlantRepository struct {
}

func InitPlantRepository() PlantRepository {
	return PlantRepository{}
}

func (p *PlantRepository) GetAllPlantsRepository() ([]entities.Plant, error) {
	var plants []entities.Plant
	if err := database.DB.Find(&plants).Error; err != nil {
		return nil, err
	}
	return plants, nil
}

func (p *PlantRepository) GetAllPlantsAvailableRepository() ([]entities.Plant, error) {
	var plants []entities.Plant
	if err := database.DB.Where("status = ?", "available").Find(&plants).Error; err != nil {
		return nil, err
	}
	return plants, nil
}

func (p *PlantRepository) CreatePlantRepository(createPlantRequest plant.CreatePlantRequest) (entities.Plant, error) {

	newPlant := model.Plant{
		PlantID:     uuid.New(),
		Name:        createPlantRequest.Name,
		Description: createPlantRequest.Description,
		Price:       createPlantRequest.Price,
		Picture:     "url", // Set picture URL here
		Status:      "available",
	}

	if err := database.DB.Create(&newPlant).Error; err != nil {
		return entities.Plant{}, err
	}

	createdPlant := entities.Plant{
		PlantID:     newPlant.PlantID,
		Name:        newPlant.Name,
		Description: newPlant.Description,
		Price:       newPlant.Price,
		Picture:     newPlant.Picture,
		Status:      newPlant.Status,
		CreatedAt:   newPlant.CreatedAt,
		UpdatedAt:   newPlant.UpdatedAt,
	}

	return createdPlant, nil
}

func (p *PlantRepository) GetPlantByIdRepository(plantId string) (entities.Plant, error) {
	var plant entities.Plant
	if err := database.DB.Where("plant_id = ?", plantId).First(&plant).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return plant, errors.New("plant not found")
		}

		return plant, err

	}
	return plant, nil
}

func (p *PlantRepository) DeletePlantByIdRepository(plantId string) error {
	result := database.DB.Where("plant_id = ?", plantId).Delete(&model.Plant{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("plant not found")
	}
	return nil
}

func (p *PlantRepository) EditPlantByIdRepository(editPlantRequest plant.EditPlantRequest, plantId string) (entities.Plant, error) {
	var plant model.Plant
	if err := database.DB.Where("plant_id = ?", plantId).First(&plant).Error; err != nil {
		return entities.Plant{}, err
	}

	plant.Name = editPlantRequest.Name
	plant.Description = editPlantRequest.Description
	plant.Status = "unavailable"
	plant.Price = editPlantRequest.Price

	if err := database.DB.Save(&plant).Error; err != nil {
		return entities.Plant{}, err
	}

	return entities.Plant{
		PlantID:     plant.PlantID,
		Name:        plant.Name,
		Description: plant.Description,
		Price:       plant.Price,
	}, nil

}
