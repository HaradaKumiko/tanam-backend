package admin

import (
	"net/http"
	"tanam-backend/domains/web/plant"
	"tanam-backend/helpers"
	"tanam-backend/helpers/response"
	"tanam-backend/services/admin"

	"github.com/labstack/echo/v4"
)

type PlantController struct {
	service admin.PlantService
}

func InitPlantController() PlantController {
	return PlantController{
		service: admin.InitPlantService(),
	}
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

	response := response.SuccessSingularFormatter("Data Semua Tanaman", plants)
	return c.JSON(http.StatusOK, response)
}

func (controller *PlantController) CreatePlantController(c echo.Context) error {
	var plant plant.CreatePlantRequest
	if err := c.Bind(&plant); err != nil {
		response := response.ErrorFormatter(err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}
	form, err := c.MultipartForm()

	if err != nil {
		response := response.ErrorFormatter(err.Error())
		return c.JSON(http.StatusInternalServerError, response)
	}

	picture := form.File["picture"]
	if len(picture) == 0 {
		response := response.ErrorFormatter("Gambar Tanaman Tidak Boleh Kosong")
		return c.JSON(http.StatusBadRequest, response)
	}

	if len(picture) > 1 {
		response := response.ErrorFormatter("Gambar Tanaman Hanya Boleh Satu")
		return c.JSON(http.StatusBadRequest, response)
	}
	for _, file := range picture {
		if !helpers.IsImageFile(file.Filename) {
			response := response.ErrorFormatter("Format file gambar tidak didukung")
			return c.JSON(http.StatusBadRequest, response)
		}
	}

	plant.Picture = picture
	data, err := controller.service.CreatePlantService(plant)

	if err != nil {
		response := response.ErrorFormatter(err.Error())
		return c.JSON(http.StatusInternalServerError, response)
	}

	return c.JSON(http.StatusOK, data)
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

func (controller *PlantController) EditPlantByPlantIdController(c echo.Context) error {
	plantId := c.Param("plantId")
	var request plant.EditPlantRequest

	if err := c.Bind(&request); err != nil {
		response := response.ErrorFormatter(err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	plant, err := controller.service.EditPlantByIdService(request, plantId)
	if err != nil {
		response := response.ErrorFormatter(err.Error())
		return c.JSON(http.StatusBadRequest, response)
	}

	response := response.SuccessSingularFormatter("Berhasil Edit Data Tanaman", plant)

	return c.JSON(http.StatusOK, response)
}

func (controller *PlantController) DeletePlantByIdController(c echo.Context) error {
	plantId := c.Param("plantId")
	err := controller.service.DeletePlantByIdService(plantId)

	if err != nil {
		response := response.ErrorFormatter(err.Error())
		return c.JSON(http.StatusNotFound, response)
	}

	response := response.SuccessSingularFormatter("Berhasil Hapus Data Tanaman", nil)
	return c.JSON(http.StatusOK, response)
}

