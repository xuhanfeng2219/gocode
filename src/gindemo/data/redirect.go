package data

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Redirect() {
	r := gin.Default()

	r.GET("/redirect", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
	})

	if err := r.Run(":10020"); err != nil {
		fmt.Println("execute wrong!")
	}
}
