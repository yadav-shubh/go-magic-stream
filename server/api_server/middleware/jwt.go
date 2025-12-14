package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/yadav-shubh/go-magic-stream/utils"
	"go.uber.org/zap"
)

type exemptedMethod struct {
	url    string
	method string
}

var exemptedPaths = []exemptedMethod{
	{url: "/api/auth/auth-info", method: "GET"},
	{url: "/api/auth/authenticate", method: "GET"},
	{url: "/api/auth/register", method: "POST"},
}

// isExempted checks if the request path is exempted from JWT validation
func isExempted(requestPath, method string) bool {
	for _, p := range exemptedPaths {
		if p.method == method {
			utils.Log.Info("Exempted path", zap.String("path", p.url), zap.String("requestPath", requestPath))
			prefix := strings.TrimSuffix(p.url, "*filepath")
			if strings.HasPrefix(requestPath, prefix) {
				return true
			}
		}
	}
	return false
}

// JWTAuth validates JWT token from Authorization header
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		if isExempted(c.Request.URL.Path, c.Request.Method) {
			// Skip JWT validation for exempted APIs
			c.Next()
			return
		}

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		// Extract token from "Bearer <token>"
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid authorization header format"})
			c.Abort()
			return
		}

		tokenString := parts[1]

		// Validate token using utility function
		claims, err := utils.ValidateAccessToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		c.Set(utils.UserIDKey, claims.UserID)
		c.Set(utils.EmailKey, claims.Email)
		c.Set(utils.RoleKey, claims.Role)
		c.Set(utils.ClaimsKey, claims)
		c.Next()
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, User-Agent, Referer")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
