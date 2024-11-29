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
	case "min":
		return "The value must be at least " + fe.Param() + " characters"
	case "max":
		return "The value must not exceed " + fe.Param() + " characters"
	case "len":
		return "The value must be exactly " + fe.Param() + " characters"
	case "numeric":
		return "The value must be a number"
	case "alpha":
		return "The value must contain only alphabetic characters"
	case "alphanum":
		return "The value must contain only alphanumeric characters"
	case "url":
		return "Invalid URL format"
	case "uuid":
		return "Invalid UUID format"
	case "eqfield":
		return "This field must be equal to " + fe.Param()
	case "nefield":
		return "This field must not be equal to " + fe.Param()
	case "gte":
		return "The value must be greater than or equal to " + fe.Param()
	case "lte":
		return "The value must be less than or equal to " + fe.Param()
	case "oneof":
		return "The value must be one of the following: " + fe.Param()
	default:
		return "Invalid value"
	}
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
