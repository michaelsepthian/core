package deleteusers

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
	userID := p.ByName("userId")

	uintUserID, _ := strconv.ParseUint(userID, 10, 32)

	response := UseCase(h.db, uintUserID)
	if response == nil {
		response = &others.BaseResponse{
			Message:    "Success Delete User",
			StatusCode: http.StatusOK,
		}
	}

	utils.SendResponse(w, *response)
}
