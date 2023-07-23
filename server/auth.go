package server

import (
	"ectary/handlers/config"
	error2 "ectary/utils/error"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"log"
)

func auth(ctx *fiber.Ctx) error {
	log.Println(ctx.GetReqHeaders()["Authorization"])
	token1 := utils.CopyString(ctx.Get("Authorization"))
	if token1 != "" {
		if validToken(token1) {
			return ctx.Next()
		}
		return ctx.Status(fiber.StatusUnauthorized).JSON(requestError{
			Code:       error2.CodeErrInvalidToken,
			Message:    error2.ErrInvalidToken.Error(),
			StatusCode: fiber.StatusUnauthorized,
		})
	}
	token2 := ctx.Query("token")
	if token2 != "" {
		if validToken(token2) {
			return ctx.Next()
		}
		return ctx.Status(fiber.StatusUnauthorized).JSON(requestError{
			Code:       error2.CodeErrInvalidToken,
			Message:    error2.ErrInvalidToken.Error(),
			StatusCode: fiber.StatusUnauthorized,
		})
	}
	return ctx.Status(fiber.StatusUnauthorized).JSON(requestError{
		Code:       error2.CodeErrUnauthorized,
		Message:    error2.ErrUnauthorized.Error(),
		StatusCode: fiber.StatusUnauthorized,
	})
}
func validToken(token string) bool {
	for _, t := range config.Tokens {
		if t == token {
			return true
		}
	}
	return false
}
