package generatetoken

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

func CreateToken(claims string, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"info": claims,
		"nbf":  time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
		"exp":  time.Now().Add(time.Hour * 24).Unix()})
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
