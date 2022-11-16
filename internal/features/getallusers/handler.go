package getallusers

import (
	"github.com/go-redis/redis/v8"
	"github.com/julienschmidt/httprouter"
	"gitlab.com/systeric/internal/chat/backend/core/internal/domain/others"
	"gitlab.com/systeric/internal/chat/backend/core/internal/features/utils"
	rediscustom "gitlab.com/systeric/internal/chat/backend/core/internal/server/redis"
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
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

func (h handler) Handle(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var client *redis.Client
	client = rediscustom.Setup()
	var responses []Response

	users, response := UseCase(h.db)

	err := rediscustom.GetObject(client, "users", &responses)
	if err != nil {
		err = rediscustom.SetObject(client, "users", users, 0)
		if err != nil {
			response = &others.BaseResponse{
				Message:    "Set Redis Failed",
				StatusCode: http.StatusBadRequest,
			}
		}
		_ = rediscustom.GetObject(client, "users", &responses)
	}

	if len(responses) == 0 {
		err = rediscustom.SetObject(client, "users", users, 0)
		if err != nil {
			response = &others.BaseResponse{
				Message:    "Set Redis Failed",
				StatusCode: http.StatusBadRequest,
			}
		}
	}

	if response == nil {
		response = &others.BaseResponse{
			Message:    "All Data Users",
			StatusCode: http.StatusFound,
			Data:       responses,
		}
	}
	utils.SendResponse(w, *response)
}
