package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RoleAuth(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRole, exists := c.Get("role")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Role not found in token"})
			c.Abort()
			return
		}

		// Cek apakah role user ada di daftar allowedRoles
		for _, role := range allowedRoles {
			if role == userRole {
				c.Next()
				return
			}
		}

		c.JSON(http.StatusForbidden, gin.H{"error": "You don't have permission to access this resource"})
		c.Abort()
	}
}
