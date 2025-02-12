package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/Jorozcob/todo-golang/app/dtos"
	"github.com/Jorozcob/todo-golang/config"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"Error": "Authorization header is required"})
			c.Abort()
			return
		}

		bearerToken := strings.Split(authHeader, " ")

		if len(bearerToken) != 2 || bearerToken[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		claims, err := validateToken(bearerToken[1])
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		c.Set("UserID", claims.UserID)
		c.Set("RoleID", claims.RoleID)
		c.Set("RoleName", claims.RoleName)
		c.Next()
	}
}

func RoleMiddleware(allowed map[string]bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("RoleName")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Role not found in context"})
			c.Abort()
			return
		}

		if !allowed[role.(string)] {
			c.JSON(http.StatusForbidden, gin.H{"error": "You don't have permission to access this resource"})
			c.Abort()
			return
		}
		c.Next()
	}
}

func validateToken(tokenString string) (*dtos.Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &dtos.Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(config.CFG.JWTSecret), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*dtos.Claims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
