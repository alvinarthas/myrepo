package middleware

import (
	"fmt"
	"net/http"

	"github.com/alvinarthas/myrepo/config"
	"github.com/alvinarthas/myrepo/utils/common"
	"github.com/alvinarthas/myrepo/utils/response"
)

func AppAuthorization(h http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		apiKeyHeader := req.Header.Get("x-api-key")
		if apiKeyHeader == "" {
			response.Error(res, common.RecordError(fmt.Errorf("undefined api key"), http.StatusForbidden, "Forbiden access"))
			return
		}
		if isAppKeyValid := validateApiKey(apiKeyHeader); !isAppKeyValid {
			response.Error(res, common.RecordError(fmt.Errorf("invalid api key"), http.StatusForbidden, "Forbiden access"))
			return
		}
		h.ServeHTTP(res, req)
	})
}

// check given key and compare with app key from env
func validateApiKey(myKey string) bool {
	apiSecretKey := config.CONFIG.Server.XApiKey
	return myKey == apiSecretKey
}
