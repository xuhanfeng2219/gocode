package session

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthMiddWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		if cookie, err := c.Cookie("abc"); err != nil {
			if cookie == "123" {
				c.Next()
				return
			}
		}
		c.JSON(http.StatusUnauthorized, gin.H{"error": "sss"})
		c.Abort()
		return
	}
}

func main() {
	r := gin.Default()
	r.GET("/login", func(c *gin.Context) {
		c.SetCookie("abc", "123", 60, "/", "localhost", false, true)
		c.String(http.StatusOK, "login success")
	})
	r.GET("/home", AuthMiddWare(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "home"})
	})
	if err := r.Run(":8000"); err != nil {
		fmt.Println("run wrong!")
	}
}
