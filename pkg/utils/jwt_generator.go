package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
)

type Tokens struct {
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
}

func GenerateNewTokens(id string) (*Tokens, error) {
	accessToken, err := generateNewAccessToken(id)
	if err != nil {
		return nil, err
	}
	refreshToken, err := generateNewRefreshToken()
	if err != nil {
		return nil, err
	}
	return &Tokens{
		Access:  accessToken,
		Refresh: refreshToken,
	}, nil
}

func generateNewAccessToken(id string) (string, error) {
	secret := viper.GetString("jwt.secret-key")
	minutesCount := viper.GetInt("jwt.secret-key-expire-minutes")
	claims := &jwt.MapClaims{
		"id":     id,
		"expire": time.Now().Add(time.Minute * time.Duration(minutesCount)).Unix(),
	}
	t, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secret))
	return t, err
}

func generateNewRefreshToken() (string, error) {
	hash := sha256.New()
	refresh := viper.GetString("jwt.refresh-key") + time.Now().String()
	_, err := hash.Write([]byte(refresh))
	if err != nil {
		return "", err
	}
	minutesCount := viper.GetInt("jwt.refresh-key-expire-minutes")
	expireTime := fmt.Sprint(time.Now().Add(time.Minute * time.Duration(minutesCount)).Unix())
	t := hex.EncodeToString(hash.Sum(nil)) + "." + expireTime
	return t, nil
}

func ParseRefreshToken(refreshToken string) (int64, error) {
	return strconv.ParseInt(strings.Split(refreshToken, ".")[1], 0, 64)
}
