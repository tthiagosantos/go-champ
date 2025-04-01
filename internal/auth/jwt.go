package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var secretKey = []byte("chave-secreta-super-segura")

func GeraTokenJWT(usuario string) (string, error) {
	claims := jwt.MapClaims{
		"sub": usuario,
		"exp": time.Now().Add(time.Hour * 2).Unix(), // expira em 2 horas
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidaTokenJWT(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrSignatureInvalid
}
