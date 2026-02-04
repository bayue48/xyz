package middleware

import (
	"xyz-multifinance/internal/model/request"
	"xyz-multifinance/internal/model/response"
	"xyz-multifinance/internal/usecase"

	"github.com/gofiber/fiber/v2"
)

func NewAuth(userUserCase *usecase.CustomerUseCase) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		token := ctx.Get("Authorization")
		// split token from bearer
		if token != "" {
			token = token[7:] // Remove "Bearer " prefix
		}
		request := &request.VerifyReq{Token: token}
		userUserCase.Log.Debugf("Authorization : %s", request.Token)

		auth, err := userUserCase.VerifyToken(ctx.UserContext(), request.Token)
		if err != nil {
			userUserCase.Log.Warnf("Failed find user by token : %+v", err)
			return fiber.ErrUnauthorized
		}

		userUserCase.Log.Debugf("User : %+v", auth.ID)
		ctx.Locals("auth", auth)
		return ctx.Next()
	}
}

func GetUser(ctx *fiber.Ctx) *response.Auth {
	return ctx.Locals("auth").(*response.Auth)
}
