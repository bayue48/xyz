package errors

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"unicode"

	"github.com/go-playground/validator/v10"
)

func NewValidationError(message string, err map[string]string) AppError {
	return AppError{
		Code:    http.StatusBadRequest, // HTTP status code for bad request
		Message: message,
		Err:     err, // Validation details
	}
}

func ValidationErrorToAppError(err error) AppError {
	validationErrors := make(map[string]string)

	var errs validator.ValidationErrors
	if errors.As(err, &errs) {
		for _, e := range errs {
			validationErrors[ToSnakeCase(e.Field())] = validationErrorToText(e)
		}
	} else {
		return NewDefaultError(http.StatusBadRequest, err.Error())
	}

	return NewValidationError("Validation failed", validationErrors)
}

var errorWording = map[string]string{
	"required":         "is required",
	"email":            "is an invalid email address",
	"is_digit":         "is not a digit",
	"is_only_alphabet": "is not a valid alphabet",
	"min":              "is too short",
	"max":              "is too long",
	"len":              "has an invalid length",
}

func validationErrorToText(e validator.FieldError) string {
	if wording, ok := errorWording[e.Tag()]; ok {
		return wording
	}
	return "invalid"
}

func ToSnakeCase(str string) string {
	var charArr []string

	for i, char := range str {
		if i > 0 && unicode.IsUpper(rune(str[i-1])) && unicode.IsUpper(rune(str[i])) {
			charArr = append(charArr, string(char))
			continue
		}

		if i > 0 && unicode.IsUpper(rune(str[i])) {
			charArr = append(charArr, "_")
		}

		if i > 1 && unicode.IsUpper(rune(str[i-1])) && unicode.IsUpper(rune(str[i-2])) && unicode.IsLower(rune(str[i])) {
			charArr, _ = InsertAtIndex(charArr, i-1, "_")
		}

		charArr = append(charArr, string(char))
	}

	return strings.ToLower(strings.Join(charArr, ""))
}

func InsertAtIndex[T any](slice []T, index int, value T) ([]T, error) {
	if index < 0 || index > len(slice) {
		return nil, fmt.Errorf("index out of range")
	}

	newSlice := make([]T, len(slice)+1)
	copy(newSlice, slice[:index])
	newSlice[index] = value
	copy(newSlice[index+1:], slice[index:])

	return newSlice, nil
}
