package middlewares

import (
	"fmt"
	"log"
	"strings"

	"api_gateway/config"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func JWTMiddlewares(ctx *gin.Context) {
	jwtKey := []byte(config.Load().SECRET_KEY)
	tokenString := ctx.GetHeader("Authorization")
	log.Println(string(jwtKey))

	if !strings.HasPrefix(tokenString, "Bearer ") {
		ctx.JSON(401, gin.H{"error": "token not found"})
		ctx.Abort()
		return
	}
	tokenString = strings.TrimPrefix(tokenString, "Bearer ")
	fmt.Println(tokenString)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return jwtKey, nil
	})
	fmt.Println(token)
	if err != nil {
		ctx.JSON(401, gin.H{"error": "parsing token failed: " + err.Error()})
		ctx.Abort()
		return
	}

	if !token.Valid {
		ctx.JSON(401, gin.H{"error": "invalid token"})
		ctx.Abort()
		return
	}

	ctx.Next()
}
