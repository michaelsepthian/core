package createusers

import (
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/julienschmidt/httprouter"
	"gitlab.com/systeric/internal/chat/backend/core/internal/domain/others"
	"gitlab.com/systeric/internal/chat/backend/core/internal/features/utils"
	rediscustom "gitlab.com/systeric/internal/chat/backend/core/internal/server/redis"
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
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

type Response struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

func (h handler) Handle(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var client *redis.Client
	client = rediscustom.Setup()

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

	user, response := UseCase(h.db, requestBody.Name, requestBody.Age, requestBody.Address)

	if response == nil {
		response = &others.BaseResponse{
			Message:    "User Created",
			StatusCode: http.StatusCreated,
			Data: Response{
				Name:    user.Name,
				Age:     user.Age,
				Address: user.Address,
			},
		}
	}

	err = rediscustom.SetObject(client, "users", response.Data, 0)
	if err != nil {
		response = &others.BaseResponse{
			Message:    "Set Redis Failed",
			StatusCode: http.StatusBadRequest,
		}
	}

	utils.SendResponse(w, *response)
}
