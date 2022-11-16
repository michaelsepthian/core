package getallorders

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

type Response struct {
	NameProduct string `json:"nameProduct"`
	Quantity    int    `json:"quantity"`
}

func (h handler) Handle(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	userID := r.URL.Query().Get("userId")

	uintUserID, _ := strconv.ParseUint(userID, 10, 32)

	orders, response := UseCase(h.db, uintUserID)
	if response == nil {
		response = &others.BaseResponse{
			Message:    "All Data Products",
			StatusCode: http.StatusFound,
			Data:       orders,
		}
	}

	utils.SendResponse(w, *response)
}
