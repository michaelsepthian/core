package updateproducts

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"gitlab.com/systeric/internal/chat/backend/core/internal/domain/others"
	"gitlab.com/systeric/internal/chat/backend/core/internal/features/utils"
	"gorm.io/gorm"
	"io"
	"net/http"
	"strconv"
)

type handler struct {
	db *gorm.DB
}

func New(db *gorm.DB) httprouter.Handle {
	return handler{db: db}.Handle
}

type Request struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
	Stock int    `json:"stock"`
}

type Response struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
	Stock int    `json:"stock"`
}

func (h handler) Handle(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	productID := p.ByName("productId")

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

	uintProductID, _ := strconv.ParseUint(productID, 10, 32)
	product, response := UseCase(h.db, uintProductID, requestBody.Name, requestBody.Price, requestBody.Stock)
	if response == nil {
		response = &others.BaseResponse{
			Message:    "Update Data Product Success",
			StatusCode: http.StatusOK,
			Data: Response{
				Name:  product.Name,
				Price: product.Price,
				Stock: product.Stock,
			},
		}
	}

	utils.SendResponse(w, *response)
}
