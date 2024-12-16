package payments

import (
	"strings"
)

type PaymentRequest struct {
	PaymentMethod *string            `json:"payment_method"`
	Amount        *float64           `json:"amount"`
	Description   *string            `json:"description"`
	Recipient     *map[string]string `json:"recipient"`
	Metadata      *map[string]string `json:"metadata"`
}

type PaymentResponse struct {
	Status       string `json:"status"`
	PaymentToken int64  `json:"payment_token"`
}

type ValidationErrors struct {
	Errors map[string]string
}

func (p PaymentRequest) Validate() error {
	validationErrors := ValidationErrors{Errors: make(map[string]string)}
	if p.Amount == nil {
		validationErrors.AddError("Amount", "Amount is required")
	}
	if p.PaymentMethod == nil || strings.TrimSpace(*p.PaymentMethod) == "" {
		validationErrors.AddError("PaymentMethod", "PaymentMethod is required")
	}
	if len(validationErrors.Errors) == 0 {
		return nil
	}
	return validationErrors
}

func (v ValidationErrors) AddError(property, message string) {
	v.Errors[property] = message
}

func (v ValidationErrors) Error() string {
	return "Validation error"
}
