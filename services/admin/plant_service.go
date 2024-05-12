package admin

import (
	"tanam-backend/domains/web/plant"
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
	plants, err := service.repository.GetAllPlantsRepository()
	if err != nil {
		return nil, err
	}
	return plants, nil
}

func (service *PlantService) CreatePlantService(createPlantRequest plant.CreatePlantRequest) (entities.Plant, error) {
	plant, err := service.repository.CreatePlantRepository(createPlantRequest)
	if err != nil {
		return plant, err
	}
	return plant, nil
}

func (service *PlantService) GetPlantByIdService(plantId string) (entities.Plant, error) {
	plant, err := service.repository.GetPlantByIdRepository(plantId)
	if err != nil {
		return plant, err
	}

	return plant, nil
}

func (service *PlantService) EditPlantByIdService(editPlantRequest plant.EditPlantRequest, plantId string) (entities.Plant, error) {
	plant, err := service.repository.EditPlantByIdRepository(editPlantRequest, plantId)
	if err != nil {
		return plant, err
	}
	return plant, nil
}

func (service *PlantService) DeletePlantByIdService(plantId string) error {

	err := service.repository.DeletePlantByIdRepository(plantId)
	if err != nil {
		return err
	}
	return nil
}
