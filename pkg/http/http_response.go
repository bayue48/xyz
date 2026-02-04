package http

import (
	"errors"
	"net/http"
	"strings"
	internal_err "xyz-multifinance/pkg/errors"

	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
	Error   any    `json:"error,omitempty"`
}

func SendSuccess(c *fiber.Ctx, status int, message string, data any) error {
	var msg string
	if message != "" {
		msg = message
	} else {
		msg = http.StatusText(status)
	}

	response := Response{
		Status:  status,
		Message: msg,
		Data:    data,
	}
	c.Status(status).JSON(response)
	return nil
}

func SendError(c *fiber.Ctx, err error) error {
	var appErr internal_err.AppError
	if errors.As(err, &appErr) {
		response := Response{
			Status:  appErr.Code,
			Message: appErr.Message,
			Error:   appErr.Err,
		}
		c.Status(appErr.Code).JSON(response)
		return nil
	}
	// Handle specific Fiber errors
	if errors.Is(err, fiber.ErrBadRequest) {
		response := Response{
			Status:  fiber.ErrBadRequest.Code,
			Message: fiber.ErrBadRequest.Message,
			Error:   fiber.ErrBadRequest.Error(),
		}
		c.Status(fiber.ErrBadRequest.Code).JSON(response)
		return nil
	}
	if errors.Is(err, fiber.ErrConflict) {
		response := Response{
			Status:  fiber.ErrConflict.Code,
			Message: fiber.ErrConflict.Message,
			Error:   fiber.ErrConflict.Error(),
		}
		c.Status(fiber.ErrConflict.Code).JSON(response)
		return nil
	}
	if errors.Is(err, fiber.ErrUnauthorized) {
		response := Response{
			Status:  fiber.ErrUnauthorized.Code,
			Message: fiber.ErrUnauthorized.Message,
			Error:   fiber.ErrUnauthorized.Error(),
		}
		c.Status(fiber.ErrUnauthorized.Code).JSON(response)
		return nil
	}
	if errors.Is(err, fiber.ErrNotFound) {
		response := Response{
			Status:  fiber.ErrNotFound.Code,
			Message: fiber.ErrNotFound.Message,
			Error:   fiber.ErrNotFound.Error(),
		}
		c.Status(fiber.ErrNotFound.Code).JSON(response)
		return nil
	}
	if strings.HasPrefix(err.Error(), "insufficient") || strings.HasPrefix(err.Error(), "customer don't") {
		response := Response{
			Status:  fiber.ErrBadRequest.Code,
			Message: fiber.ErrBadRequest.Message,
			Error:   err.Error(),
		}
		c.Status(fiber.ErrBadRequest.Code).JSON(response)
		return nil
	}
	// Fallback for unknown errors
	response := Response{
		Status:  http.StatusInternalServerError,
		Message: http.StatusText(http.StatusInternalServerError),
		Error:   err.Error(),
	}
	c.Status(http.StatusInternalServerError).JSON(response)
	return nil
}
