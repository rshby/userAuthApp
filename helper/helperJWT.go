package helper

import (
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
	"userAuthApp/model/auth"
)

func GenerateToken(accountId int) (string, error) {
	claims := &auth.JwtClaims{
		Id: accountId,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "authAPP",
			ExpiresAt: jwt.NewNumericDate(time.Now().Local().Add(1 * time.Hour)),
		}}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := tokenClaims.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}

	// sucess generate token
	return tokenString, nil
}
