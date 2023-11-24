package utils

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type JWT struct {
	Claims		jwt.MapClaims
	Key			[]byte
	TokenStr	string
}

func (j *JWT) GenerateToken() string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, j.Claims)
	signedToken, _ := token.SignedString(j.Key)

	return signedToken
}

func (j *JWT) VerifyToken() (*jwt.MapClaims, error) {
	token, err := jwt.ParseWithClaims(j.TokenStr, &jwt.MapClaims{}, func (token *jwt.Token) (interface{}, error) {
		return j.Key, nil
	})

	if err != nil {
		return nil, err
	}

	if token.Valid {
		return token.Claims.(*jwt.MapClaims), nil
	}

	return nil, fmt.Errorf("token not valid")
}