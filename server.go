package audiotranscribe

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func SendError(w http.ResponseWriter, err error, code int) {
	bqResp := new(BigQueryResponse)
	bqResp.ErrorMessage = fmt.Sprintf("Got error with details: %v", err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(bqResp)
}

func SendSuccess(w http.ResponseWriter, bqResp *BigQueryResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bqResp)
}
