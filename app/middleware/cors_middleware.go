package middleware

import (
	"strings"

	"github.com/Jorozcob/todo-golang/config"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func GetCorsMiddleware() gin.HandlerFunc {
	corsConfig := cors.DefaultConfig()
	allowedoriginsString := config.CFG.AllowedOrigins
	allowedOrigins := strings.Split(allowedoriginsString, ",")
	corsConfig.AllowOrigins = allowedOrigins
	return cors.New(corsConfig)
}
