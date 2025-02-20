package security

import (
	"go-fiber-template/pkg/config"
	"time"

	"github.com/golang-jwt/jwt"
)

const (
	jwtExpIn = time.Hour * 168
)

func GenerateJWTToken(payload map[string]interface{}, cfg *config.Config) (string, error) {
	claims := jwt.MapClaims{}
	expIn := time.Now().Add(jwtExpIn).Unix()
	claims["exp"] = expIn
	claims["iat"] = time.Now().Unix()
	claims["iss"] = cfg.TokenIssuer
	for key, value := range payload {
		claims[key] = value
	}


	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(cfg.SecretKey))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
