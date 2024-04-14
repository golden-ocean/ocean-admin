package utils

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

type TokenMetadata struct {
	ID     string
	Expire int64
}

func ExtractTokenMetadata(c *fiber.Ctx) (*TokenMetadata, error) {
	token, err := verifyToken(c)
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return &TokenMetadata{
			ID:     claims["id"].(string),
			Expire: int64(claims["expire"].(float64)),
		}, nil
	}
	return nil, err
}

func extractToken(c *fiber.Ctx) string {
	bearToken := c.Get("Authorization")
	onlyToken := strings.Split(bearToken, " ")
	if len(onlyToken) == 2 {
		return onlyToken[1]
	}
	return ""
}

func verifyToken(c *fiber.Ctx) (*jwt.Token, error) {
	tokenString := extractToken(c)
	token, err := jwt.Parse(tokenString, jwtKeyFunc)
	if err != nil {
		return nil, err
	}
	return token, nil
}

func jwtKeyFunc(token *jwt.Token) (interface{}, error) {
	return []byte(viper.GetString("jwt.secret-key")), nil
}
