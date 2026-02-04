package http

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func BindRequest(ctx *fiber.Ctx, v *validator.Validate, reqStruct any) error {
	if err := ctx.BodyParser(reqStruct); err != nil {
		logrus.Warnf("Failed to parse request body: %v", err)
		return fiber.ErrBadRequest
	}
	// SanitizeStruct(reqStruct)
	if err := v.Struct(reqStruct); err != nil {
		logrus.Warnf("Validation failed for request body: %v", err)
		return fiber.ErrBadRequest
	}
	return nil
}

func SanitizeStruct(payload any) {
	v := reflect.ValueOf(payload).Elem()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		if field.Kind() == reflect.String {
			sanitizedValue := strings.TrimSpace(field.String())
			field.SetString(sanitizedValue)
		}
	}
}
