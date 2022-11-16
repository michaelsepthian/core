package cancelorder

import (
	"github.com/julienschmidt/httprouter"
	"gitlab.com/systeric/internal/chat/backend/core/internal/domain/others"
	"gitlab.com/systeric/internal/chat/backend/core/internal/features/utils"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type handler struct {
	db *gorm.DB
}

func New(db *gorm.DB) httprouter.Handle {
	return handler{db: db}.Handle
}

func (h handler) Handle(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	orderID := r.URL.Query().Get("orderId")
	userID := r.URL.Query().Get("userId")

	uintOrderID, _ := strconv.ParseUint(orderID, 10, 32)
	uintUserID, _ := strconv.ParseUint(userID, 10, 32)

	response := UseCase(h.db, uintOrderID, uintUserID)
	if response == nil {
		response = &others.BaseResponse{
			Message:    "Success Delete Order",
			StatusCode: http.StatusOK,
		}
	}

	utils.SendResponse(w, *response)
}
