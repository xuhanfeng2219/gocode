package data

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func Async() {
	r := gin.Default()

	r.GET("/async", func(c *gin.Context) {
		cContext := c.Copy()

		go func() {
			time.Sleep(3 * time.Second)
			log.Printf("异步执行：%s", cContext.Request.URL.Path)
		}()
	})

	r.GET("/sync", func(c *gin.Context) {
		time.Sleep(3 * time.Second)
		log.Printf("同步执行:%s", c.Request.URL.Path)
	})

	if err := r.Run(":9000"); err != nil {
		fmt.Println("some thing is executed wrong!")
	}

}
