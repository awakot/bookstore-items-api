package formatter

import (
	"encoding/json"
	"github.com/waytkheming/bookstore-items-api/utils/errors"
	"net/http"
)

func RespondJSON(w http.ResponseWriter, statusCode int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(body)
}

func RespondErr(w http.ResponseWriter, err errors.RestErr) {
	RespondJSON(w, err.Code, err)
}
