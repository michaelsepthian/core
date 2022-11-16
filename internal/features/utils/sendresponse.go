package utils

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"gitlab.com/systeric/internal/chat/backend/core/internal/domain/others"
	"net/http"
)

func SendResponse(w http.ResponseWriter, response others.BaseResponse) {
	responseJSON, err := json.Marshal(response)
	if err != nil {
		SendResponse(w, others.BaseResponse{
			Message:    "Error Create Response",
			StatusCode: http.StatusInternalServerError,
		})
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)
	_, err = w.Write(responseJSON)
	if err != nil {
		log.Error().Err(err).Msg("Write Response failed")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
