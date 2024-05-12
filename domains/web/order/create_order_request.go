package order

type CreateOrderRequest struct {
	PlantID  string `json:"plant_id"`
	Quantity int32  `json:"quantity"`
}

type CustomerDetail struct {
	FName string
	Email string
	Phone string
}

type ItemDetails struct {
	ID    string
	Name  string
	Price int64
	Qty   int32
}

type CreateOrderMidtransRequest struct {
	PlantID        string
	OrderID        string
	Price          int64
	CustomerDetail CustomerDetail
	ItemDetails    ItemDetails
}
