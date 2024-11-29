package resources

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

// Struct untuk pesan error
type ErrorMsg struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// Fungsi untuk mengembalikan pesan error berdasarkan tag validator
func GetErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Invalid email format"
	}
	return "Unknown error"
}

// Fungsi untuk memproses semua error validasi
func ProcessValidationErrors(err error) []ErrorMsg {
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		out := make([]ErrorMsg, len(ve))
		for i, fe := range ve {
			out[i] = ErrorMsg{Field: fe.Field(), Message: GetErrorMsg(fe)}
		}
		return out
	}
	return nil
}
