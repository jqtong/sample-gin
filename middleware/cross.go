package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Cross check Cross-domain problem
func Cross() gin.HandlerFunc {

	return func(c *gin.Context) {

		origin := c.Request.Header.Get("origin")
		domain := "*"
		if strings.Contains(origin, "localhost") {
			domain = origin
		}
		c.Writer.Header().Set("Access-Control-Allow-Origin", domain)
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		c.Next()
	}
}
