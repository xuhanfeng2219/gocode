package log

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

func HandleLog() {
	gin.DisableConsoleColor()

	f, _ := os.Create("gin.log")

	gin.DefaultWriter = io.MultiWriter(f)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	if err := r.Run(); err != nil {
		fmt.Println("run wrong!")
	}
}
