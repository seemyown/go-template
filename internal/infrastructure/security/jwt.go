package security

import (
	"os/user"
	"time"

	"github.com/golang-jwt/jwt"
)

func issure() string {
	u, err := user.Current()
	if err != nil {
		return "jwt.issure"
	}
	return u.Username
}

var defaultPayload = map[string]interface{}{
	"exp": time.Now().Add(time.Hour * 168).Unix(),
	"iat": time.Now().Unix(),
	"nbf": time.Now().Unix(),
	"iss": issure(),
}

func GenerateJWTToken(payload map[string]interface{}, secretKey string) (string, error) {
	claims := jwt.MapClaims{}

	for k, v := range defaultPayload {
		claims[k] = v
	}

	for key, value := range payload {
		claims[key] = value
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
