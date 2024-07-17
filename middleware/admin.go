package middleware

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IsAdmin(c *gin.Context) {
	// Get the user information from the context
	user, err := c.Get("user")
	if err {
		c.JSON(http.StatusUnauthorized, gin.H{"error": true, "message": "Unauthorized"})
		c.Abort()
		return
	}
	fmt.Println(user, "FFFFFF")
	// Check if the user is an admin
	if user != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": true, "message": "Unauthorized"})
		c.Abort()
		return
	}
	// Continue to the next middleware
	c.Next()
}
