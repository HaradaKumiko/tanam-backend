package controllers

import (
	"context"
	"net/http"
	"path/filepath"
	"tanam-backend/helpers/response"
	"tanam-backend/services"
	"tanam-backend/utils/google"

	"github.com/labstack/echo/v4"
)

type PlantController struct {
	service services.PlantService
}

func InitPlantController() PlantController {
	return PlantController{
		service: services.InitPlantService(),
	}
}

func (controller *PlantController) UploadFileController(c echo.Context) error {

	// Get the file from the form data
	file, fileHeader, err := c.Request().FormFile("picture")
	if err != nil {
		response := response.ErrorFormatter(err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}
	defer file.Close()

	// Get the file extension
	ext := filepath.Ext(fileHeader.Filename)

	// Initialize the Google Cloud Storage uploader
	ctx := context.Background()
	// Adjust the object name as needed (e.g., using a unique filename)
	objectName := "unique_filename" + ext
	url, err := google.Upload.UploadFile(ctx, file, objectName) // Use google.Upload to call the UploadFile method
	if err != nil {
		response := response.ErrorFormatter(err.Error())
		return c.JSON(http.StatusInternalServerError, response)
	}

	// Return the URL of the uploaded file
	return c.JSON(http.StatusOK, map[string]string{"url": url})

}

func (controller *PlantController) GetAllPlantController(c echo.Context) error {
	plants, err := controller.service.GetAllPlantsService()

	if err != nil {
		response := response.ErrorFormatter(err.Error())
		return c.JSON(http.StatusInternalServerError, response)
	}

	if len(plants) < 1 {
		response := response.ErrorFormatter("Data Tanaman Tidak Ditemukan")
		return c.JSON(http.StatusOK, response)
	}

	response := response.SuccessPluralFormatter("Data Semua Tanaman", plants)
	return c.JSON(http.StatusOK, response)
}

func (controller *PlantController) GetPlantByPlantIdController(c echo.Context) error {
	plantId := c.Param("plantId")
	plant, err := controller.service.GetPlantByIdService(plantId)

	if err != nil {
		response := response.ErrorFormatter(err.Error())
		return c.JSON(http.StatusNotFound, response)
	}

	response := response.SuccessSingularFormatter("Data Tanaman", plant)
	return c.JSON(http.StatusOK, response)
}
