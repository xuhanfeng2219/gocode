package data

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ParserHTML() {
	r := gin.Default()
	r.LoadHTMLGlob("./template/index.html")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{"title": "topic", "ce": "123456"})
	})

	if err := r.Run(":8000"); err != nil {
		fmt.Println("execute test wrong!")
	}
}
