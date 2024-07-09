package middlewares

import (
	"fmt"
	"github.com/Projects/Restaurant_Reservation_System/api_gateway/config"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)




func JWTMiddlewares(ctx *gin.Context) {
	jwtKey := []byte(config.Load().SECRET_KEY)
	tokenString := ctx.GetHeader("Authorization")

	if !strings.HasPrefix(tokenString, "Bearer ") {
		ctx.JSON(401, gin.H{"error": "token not found"})
		ctx.Abort()
		return
	}

	tokenString = strings.TrimPrefix(tokenString, "Berer ")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return jwtKey, nil
	})

	if err != nil || !token.Valid {
		ctx.JSON(401, gin.H{"error": "invalid token"})
		ctx.Abort()
		return
	}


	ctx.Next()
}
