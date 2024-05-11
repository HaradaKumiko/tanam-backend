package plant

type EditPlantRequest struct {
	Name        string  `json:"name" form:"name"`
	Description string  `json:"description" form:"description"`
	Status      string  `json:"status"`
	Price       float64 `json:"price" form:"price"`
}
