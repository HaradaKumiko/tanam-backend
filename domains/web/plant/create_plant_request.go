package plant

import (
	"mime/multipart"
)

type CreatePlantRequest struct {
	Name        string                  `json:"name" form:"name"`
	Description string                  `json:"description" form:"description"`
	PictureFile []*multipart.FileHeader `json:"picture_file" form:"picture_file"`
	Picture     string                  `gorm:"type:varchar(255);not null"`
	Price       float64                 `json:"price" form:"price"`
}
