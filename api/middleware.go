package api

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func (s *Server) AuthMiddleware() gin.HandlerFunc {
    return func(ctx *gin.Context) {
        auth := ctx.GetHeader("Authorization")
        if auth == "" || !strings.HasPrefix(auth, "Bearer ") {
            ctx.JSON(http.StatusUnauthorized, gin.H{"error": "missing token"})
            ctx.Abort()
            return
        }

        token := strings.TrimPrefix(auth, "Bearer ")

        parsed, err := s.tokenMaker.VerifyToken(token)
        if err != nil || !parsed.Valid {
            ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
            ctx.Abort()
            return
        }

        claims, ok := parsed.Claims.(jwt.MapClaims)
        if !ok {
            ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token claims"})
            ctx.Abort()
            return
        }
        ctx.Set("user_id", int64(claims["user_id"].(float64)))

        ctx.Next()
    }
}
