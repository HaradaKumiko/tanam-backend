package admin

import (
	"context"
	"mime/multipart"
	"path/filepath"
	"tanam-backend/domains/web/plant"
	"tanam-backend/entities"
	"tanam-backend/repository"
	"tanam-backend/utils/google"

	"github.com/google/uuid"
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

func (service *PlantService) CreatePlantService(createPlantRequest plant.CreatePlantRequest, picture []*multipart.FileHeader) (entities.Plant, error) {
	file, err := picture[0].Open()
	if err != nil {
		return entities.Plant{}, err
	}
	defer file.Close()

	ext := filepath.Ext(picture[0].Filename)

	ctx := context.Background()

	objectName := uuid.NewString() + ext
	url, err := google.Upload.UploadFile(ctx, file, objectName)
	if err != nil {
		return entities.Plant{}, err
	}

	createPlantRequest.Picture = url
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

func (service *PlantService) DeletePlantByIdService(plantId string) (entities.Plant, error) {
	ctx := context.Background()

	plant, err := service.repository.GetPlantByIdRepository(plantId)
	if err != nil {
		return plant, err
	}

	objectName := filepath.Base(plant.Picture)
	err = google.Upload.DeleteFileFromGCS(ctx, objectName)
	if err != nil {
		return entities.Plant{}, err
	}

	err = service.repository.DeletePlantByIdRepository(plantId)
	if err != nil {
		return entities.Plant{}, err
	}

	return entities.Plant{}, nil
}
