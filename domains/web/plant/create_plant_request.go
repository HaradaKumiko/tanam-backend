package plant

import (
	"mime/multipart"
)

type CreatePlantRequest struct {
	Name        string                  `json:"name" form:"name"`
	Description string                  `json:"description" form:"description"`
	Picture     []*multipart.FileHeader `json:"picture" form:"picture"`
	Price       float64                 `json:"price" form:"price"`
}
