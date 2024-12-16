package payments

import (
	"encoding/json"
	"net/http"
)

type PaymentController struct {
	repo TransactionRepository
}

func NewPaymentController(repo TransactionRepository) *PaymentController {
	return &PaymentController{repo}
}

func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, map[string]string{"error": message})
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
