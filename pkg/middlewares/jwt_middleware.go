package middlewares

import (
	"strings"

	jwtMiddleware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

// JWTProtected See: https://github.com/gofiber/jwt
func JWTProtected() func(*fiber.Ctx) error {
	config := jwtMiddleware.Config{
		SigningKey:   jwtMiddleware.SigningKey{Key: []byte(viper.GetString("jwt.secret-key"))},
		ContextKey:   "jwt",
		ErrorHandler: jwtError, // token无效后执行
		// SuccessHandler: jwtSuccess, // token有效后执行
	}
	return jwtMiddleware.New(config)
}

func jwtError(c *fiber.Ctx, err error) error {
	if strings.Contains(err.Error(), "missing or malformed JWT") {
		return fiber.NewError(fiber.StatusBadRequest, "令牌缺失或错误！")
	}
	if strings.Contains(err.Error(), "token is expired") {
		return fiber.NewError(fiber.StatusBadRequest, "令牌已过期！")
	}
	return fiber.NewError(fiber.StatusUnauthorized, err.Error())
}
