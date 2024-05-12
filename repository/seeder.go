package repository

import (
	"net/http"
	"time"

	"tanam-backend/database"
	"tanam-backend/domains/model"
	"tanam-backend/helpers"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func SeedHandler(c echo.Context) error {
	// Define sample plant data
	plants := []model.Plant{
		{
			PlantID:     uuid.MustParse("e66eaa6f-7f7a-481e-901a-752635b88b06"),
			Name:        "Pohon Bambu",
			Description: "Pohon bambu merupakan tumbuhan berkayu yang memiliki batang yang keras dan seratnya kuat. Tumbuhan ini dapat tumbuh dengan cepat dan cocok untuk program reboisasi",
			Picture:     "https://storage.googleapis.com/tanam-apps/bambu.jpg",
			Price:       15_000,
			Status:      "available",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			PlantID:     uuid.MustParse("17cc2907-e016-44cc-9ae7-686227842594"),
			Name:        "Pohon Mahoni",
			Description: "Pohon mahoni adalah jenis pohon berdaun gugur yang berasal dari Amerika Tengah dan Selatan. Kayunya sering digunakan dalam industri kayu karena keindahan dan kekuatannya",
			Picture:     "https://storage.googleapis.com/tanam-apps/mahoni.jpg",
			Price:       40_000,
			Status:      "available",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			PlantID:     uuid.MustParse("bbb9ebfa-3d3c-4195-8760-3b921ba2b27b"),
			Name:        "Pohon Jati",
			Description: "Pohon jati merupakan salah satu jenis pohon kayu yang banyak dimanfaatkan dalam industri perkayuan. Tumbuhan ini tahan terhadap serangan hama dan penyakit",
			Picture:     "https://storage.googleapis.com/tanam-apps/jati.jpg",
			Price:       75_000,
			Status:      "available",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			PlantID:     uuid.MustParse("a76af044-95c4-411a-9435-ae54990f0098"),
			Name:        "Pohon Sakura",
			Description: "Pohon sakura, juga dikenal sebagai ceri Jepang, adalah salah satu simbol musim semi di Jepang. Bunga-bunga berwarna merah muda yang indah mekar setiap musim semi",
			Picture:     "https://storage.googleapis.com/tanam-apps/sakura.jpg",
			Price:       80_000,
			Status:      "available",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			PlantID:     uuid.MustParse("76d5126f-6179-4583-846f-96a3144bc740"),
			Name:        "Pohon Gaharu",
			Description: "Pohon gaharu dikenal karena kayunya yang beraroma harum dan sering digunakan dalam pembuatan minyak gaharu. Tumbuhan ini umumnya ditemukan di hutan tropis",
			Picture:     "https://storage.googleapis.com/tanam-apps/gaharu.jpg",
			Price:       65_000,
			Status:      "available",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			PlantID:     uuid.MustParse("7a24b815-accd-4e54-84db-e2acbea80da0"),
			Name:        "Pohon Beringin",
			Description: "Pohon beringin adalah jenis pohon besar yang sering dianggap sebagai simbol kekuatan dan ketahanan. Tumbuhan ini memiliki daun hijau lebat dan akar udara yang unik",
			Picture:     "https://storage.googleapis.com/tanam-apps/beringin.jpg",
			Price:       15_000,
			Status:      "unavailable",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			PlantID:     uuid.MustParse("a8a189d9-1ff3-4fd3-b098-2a27687b8b24"),
			Name:        "Pohon Pinus",
			Description: "Pohon pinus adalah jenis pohon konifera yang umumnya memiliki jarum hijau yang tajam. Kayunya sering digunakan dalam pembuatan furnitur dan konstruksi",
			Picture:     "https://storage.googleapis.com/tanam-apps/pinus.jpg",
			Price:       15_000,
			Status:      "available",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	auths := []model.Auth{
		{
			AuthID:    uuid.MustParse("58857a4a-0be9-4fa5-a8e5-a32ed50efbb9"),
			Email:     "admin.tanam@gmail.com",
			Password:  helpers.HashPassword("password123"),
			Profile:   "https://avatars.githubusercontent.com/u/42530587",
			Role:      model.RoleAdmin,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			AuthID:    uuid.MustParse("d0f348c2-c467-46c4-8068-27658add158d"),
			Email:     "fujikawachiai@gmail.com",
			Password:  helpers.HashPassword("password123"),
			Profile:   "https://i.pinimg.com/564x/a5/04/ae/a504aeca508bcf0258ca44b625750b82.jpg",
			Role:      model.RoleDonor,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	biodatas := []model.Biodata{
		{
			BiodataID:   uuid.MustParse("07577f67-3e05-4a33-b37c-bf343e46ea44"),
			Fullname:    "Fujikawa Chiai",
			PhoneNumber: "089601230456",
			AuthID:      uuid.MustParse("d0f348c2-c467-46c4-8068-27658add158d"),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	for _, plant := range plants {
		if err := database.DB.Create(&plant).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": err.Error(),
			})
		}
	}

	for _, auth := range auths {
		if err := database.DB.Create(&auth).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": err.Error(),
			})
		}
	}

	for _, biodata := range biodatas {
		if err := database.DB.Create(&biodata).Error; err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
				"error": err.Error(),
			})
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Seeder executed successfully",
	})
}
