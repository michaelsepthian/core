package createpayment

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"gitlab.com/systeric/internal/chat/backend/core/internal/domain/others"
	"gitlab.com/systeric/internal/chat/backend/core/internal/features/utils"
	"gorm.io/gorm"
	"io"
	"net/http"
)

type handler struct {
	db *gorm.DB
}

func New(db *gorm.DB) httprouter.Handle {
	return handler{db: db}.Handle
}

type Request struct {
	OrderID uint `json:"orderId"`
	UserID  uint `json:"userId"`
}

type Response struct {
	NameProduct string `json:"nameProduct"`
	Total       int    `json:"total"`
	Quantity    int    `json:"quantity"`
}

func (h handler) Handle(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	jsonBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var requestBody Request
	err = json.Unmarshal(jsonBody, &requestBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product, payment, response := UseCase(h.db, requestBody.OrderID, requestBody.UserID)
	if response == nil {
		response = &others.BaseResponse{
			Message:    "Payment Created",
			StatusCode: http.StatusCreated,
			Data: Response{
				NameProduct: product.Name,
				Total:       payment.Total,
				Quantity:    payment.Order.Quantity,
			},
		}
	}

	utils.SendResponse(w, *response)
}
