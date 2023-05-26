package utils

import (
	"fmt"
	"joglo-dev/config/app_config"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(claims *jwt.MapClaims) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	webtoken, err := token.SignedString([]byte(app_config.SECRET_KEY))

	if err != nil {
		return "", err
	}

	return webtoken, nil
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	tokenJwt, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		_, isValid := t.Method.(*jwt.SigningMethodHMAC)

		if !isValid {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(app_config.SECRET_KEY), nil
	})
	if err != nil {
		return nil, err
	}

	return tokenJwt, err
}

func DecodeToken(tokenString string) (jwt.MapClaims, error) {
	token, err := VerifyToken(tokenString)

	if err != nil {
		return nil, err
	}

	claims, isOk := token.Claims.(jwt.MapClaims)

	if isOk && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")

}
