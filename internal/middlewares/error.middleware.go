package middlewares

import "github.com/gofiber/fiber/v2"

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	return ctx.JSON(map[string]string{"error": err.Error()})
}
