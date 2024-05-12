package services

import (
	"tanam-backend/entities"
	"tanam-backend/repository"
)

type PlantService struct {
	repository repository.PlantRepository
}

func InitPlantService() PlantService {
	return PlantService{
		repository: repository.InitPlantRepository(),
	}
}

func (service *PlantService) GetAllPlantsService() ([]entities.Plant, error) {
	plants, err := service.repository.GetAllPlantsAvailableRepository()
	if err != nil {
		return nil, err
	}
	return plants, nil
}

func (service *PlantService) GetPlantByIdService(plantId string) (entities.Plant, error) {
	plant, err := service.repository.GetPlantByIdRepository(plantId)
	if err != nil {
		return plant, err
	}

	return plant, nil
}
