package midtrans

import (
	"tanam-backend/domains/web/order"
	"tanam-backend/helpers"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type MidtransPayment struct {
	Snap *snap.Client
}

func NewMidtransPayment() *MidtransPayment {
	snap := &snap.Client{}
	snap.New(helpers.GetConfig("SERVER_KEY"), midtrans.Sandbox)

	return &MidtransPayment{
		Snap: snap,
	}
}

func (mp *MidtransPayment) CreateTransaction(request order.CreateOrderMidtransRequest) (*snap.Response, error) {
	transactionDetails := midtrans.TransactionDetails{
		OrderID:  request.OrderID,
		GrossAmt: request.Price,
	}

	customerDetail := &midtrans.CustomerDetails{
		FName: request.CustomerDetail.FName,
		Email: request.CustomerDetail.Email,
		Phone: request.CustomerDetail.Phone,
	}

	itemDetails := []midtrans.ItemDetails{{
		ID:    request.ItemDetails.ID,
		Name:  request.ItemDetails.Name,
		Price: request.ItemDetails.Price,
		Qty:   request.ItemDetails.Qty,
	}}

	snapRequest := &snap.Request{
		TransactionDetails: transactionDetails,
		CustomerDetail:     customerDetail,
		Items:              &itemDetails,
	}

	return mp.Snap.CreateTransaction(snapRequest)
}
