package createorder

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"gitlab.com/systeric/internal/chat/backend/core/internal/domain/models"
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
	UserID    uint `json:"userId"`
	ProductID uint `json:"productId"`
	Quantity  int  `json:"quantity"`
}

type Response struct {
	User     models.Users    `json:"user"`
	Product  models.Products `json:"product"`
	Quantity int             `json:"quantity"`
	Status   string          `json:"status"`
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

	order, response := UseCase(h.db, requestBody.UserID, requestBody.ProductID, requestBody.Quantity)
	if response == nil {
		response = &others.BaseResponse{
			Message:    "Order Saved",
			StatusCode: http.StatusCreated,
			Data: Response{
				User:     order.User,
				Product:  order.Product,
				Quantity: order.Quantity,
				Status:   order.Status,
			},
		}
	}

	utils.SendResponse(w, *response)
}
