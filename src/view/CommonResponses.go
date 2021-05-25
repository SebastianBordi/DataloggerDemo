package view

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/sebastianbordi/DataloggerDemo/model"
)

func BasicResponse(w *http.ResponseWriter, statusCode int, message string) {
	writer := *w
	writer.WriteHeader(statusCode)
	response := &model.BasicResponse{
		StatusCode: statusCode,
		Message:    message,
		Timestamp:  time.Now(),
	}
	writer.Header().Add("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(response)
}
