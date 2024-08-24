package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/glinboy/gin-jwt-secured-demo/config"
	"github.com/glinboy/gin-jwt-secured-demo/repository"
	"github.com/golang-jwt/jwt/v5"
)

func RequireAuth(ctx *gin.Context) {
	tokenString, err := ctx.Cookie("Authorization")
	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}

	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error){
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("SECURE_SECRET"), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}

		user, err := repository.NewUserRepositoryImpl(config.DatabaseConnection()).FindByEmail(claims["sub"].(string))
		if err != nil && user.Id == 0 {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}

		ctx.Set("user", user)

		ctx.Next()
	}
}
