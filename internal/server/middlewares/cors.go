package middlewares

import (
	"net/http"

	"github.com/rs/cors"
)

var corsSetting = cors.AllowAll()

func CORSHandler(handler http.Handler) http.Handler {
	return corsSetting.Handler(handler)
}
