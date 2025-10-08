package platform_httpHelper

import (
	"encoding/json"
	"net/http"
)

func WriteJsonResponse(writer http.ResponseWriter, statusCode int, res interface{}) error {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(statusCode)
	err := json.NewEncoder(writer).Encode(res)

	return err
}
