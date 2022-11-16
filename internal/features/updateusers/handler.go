package updateusers

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
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

type Response struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

func (h handler) Handle(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	userID := p.ByName("userId")

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

	uintUserID, _ := strconv.ParseUint(userID, 10, 32)
	user, response := UseCase(h.db, uintUserID, requestBody.Name, requestBody.Age, requestBody.Address)
	if response == nil {
		response = &others.BaseResponse{
			Message:    "Update Data User Success",
			StatusCode: http.StatusOK,
			Data: Response{
				Name:    user.Name,
				Age:     user.Age,
				Address: user.Address,
			},
		}
	}

	utils.SendResponse(w, *response)
}
