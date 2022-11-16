package finishtransaction

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
	Method string `json:"method"`
}

type Response struct {
	NameProduct   string `json:"nameProduct"`
	Quantity      int    `json:"quantity"`
	Method        string `json:"method"`
	StatusPayment string `json:"statusPayment"`
	Total         int    `json:"total"`
}

func (h handler) Handle(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	transactionID := p.ByName("transactionId")

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

	uintTransactionID, err := strconv.ParseUint(transactionID, 10, 32)
	resultResponse, response := UseCase(h.db, uintTransactionID, requestBody.Method)
	if response == nil {
		response = &others.BaseResponse{
			Message:    "Transaction Success",
			StatusCode: http.StatusCreated,
			Data:       resultResponse,
		}
	}

	utils.SendResponse(w, *response)
}
