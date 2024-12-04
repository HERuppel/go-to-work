package authentication

import (
	"errors"
	"fmt"
	"go-to-work/internal/config"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(id uint64) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissions["id"] = id

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)

	return token.SignedString([]byte(config.SecretKey))
}

func ValidateToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, getVerificationKey)
	if err != nil {
		return err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}

	return errors.New("INVALID_TOKEN")
}

func ExtractUserId(authHeader string) (uint64, error) {
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	token, err := jwt.Parse(tokenString, getVerificationKey)
	if err != nil {
		return 0, err
	}

	if permissions, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		id, err := strconv.ParseUint(fmt.Sprintf("%.0f", permissions["id"]), 10, 64)
		if err != nil {
			return 0, err
		}

		return id, nil
	}

	return 0, errors.New("INVALID_TOKEN")
}

func getVerificationKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("UNEXPECTED_SIGNATURE_METHOD! %v", token.Header["alg"])
	}

	return config.SecretKey, nil
}
