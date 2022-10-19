package middware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func MiddWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("middware starting to run")

		c.Set("request", "middware")
		//execute function
		c.Next()

		status := c.Writer.Status()
		fmt.Println("middware execute end", status)

		t2 := time.Since(t)
		fmt.Println("time:", t2)
	}
}

func main() {
	r := gin.Default()
	r.Use(MiddWare())

	{
		r.GET("/ce", func(c *gin.Context) {
			req, _ := c.Get("request")
			fmt.Println("request", req)
			c.JSON(http.StatusOK, gin.H{"request": req})
		})
	}

	if err := r.Run(":5000"); err != nil {
		fmt.Println("execute wrong!")
	}
}
