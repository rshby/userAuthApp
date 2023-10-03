package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"os"
	"strings"
	"time"
	"userAuthApp/model/auth"
	"userAuthApp/model/dto"
)

type AuthMiddleware struct {
}

// function provider
func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

func (a *AuthMiddleware) Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get token from headers
		header := c.GetHeader("Authorization")
		tokenString := strings.Split(header, " ")

		if header == "" || tokenString[len(tokenString)-1] == "" {
			response := dto.ApiMessage{
				StatusCode: http.StatusUnauthorized,
				Status:     "unauthorized",
				Message:    "token not set",
			}

			c.JSON(http.StatusUnauthorized, &response)
			c.Abort()
			return
		}

		// decode claims
		claims := auth.JwtClaims{}
		_, err := jwt.ParseWithClaims(tokenString[len(tokenString)-1], &claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET_KEY")), nil
		})

		// jika token tidak valid
		if err != nil {
			response := dto.ApiMessage{
				StatusCode: http.StatusUnauthorized,
				Status:     "unauthorized",
				Message:    err.Error(),
			}
			c.JSON(http.StatusUnauthorized, &response)
			c.Abort()
			return
		}

		// expired
		if time.Now().Local().After(time.UnixMicro(claims.RegisteredClaims.ExpiresAt.Unix() * 1000000)) {
			response := dto.ApiMessage{
				StatusCode: http.StatusUnauthorized,
				Status:     "unauthorized",
				Message:    "token expired",
			}

			c.JSON(http.StatusUnauthorized, &response)
			c.Abort()
			return
		}

		// succes auth
		c.Set("username", claims.UserName)
		c.Next()
	}
}
