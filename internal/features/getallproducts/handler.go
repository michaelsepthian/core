package getallproducts

import (
	"github.com/julienschmidt/httprouter"
	"gitlab.com/systeric/internal/chat/backend/core/internal/domain/others"
	"gitlab.com/systeric/internal/chat/backend/core/internal/features/utils"
	"gorm.io/gorm"
	"net/http"
)

type handler struct {
	db *gorm.DB
}

func New(db *gorm.DB) httprouter.Handle {
	return handler{db: db}.Handle
}

type Response struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
	Stock int    `json:"stock"`
}

func (h handler) Handle(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	products, response := UseCase(h.db)
	if response == nil {
		response = &others.BaseResponse{
			Message:    "All Data Products",
			StatusCode: http.StatusFound,
			Data:       products,
		}
	}

	utils.SendResponse(w, *response)
}
