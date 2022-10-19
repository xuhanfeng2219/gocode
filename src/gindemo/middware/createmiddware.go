package middware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func MyTime(c *gin.Context) {
	start := time.Now()
	c.Next()

	since := time.Since(start)
	fmt.Println("programmer execute time is ", since)
}

func ShopIdxHandle(c *gin.Context) {
	time.Sleep(5 * time.Second)
}

func ShopHomeHandle(c *gin.Context) {
	time.Sleep(3 * time.Second)
}

func main() {
	r := gin.Default()
	r.Use(MyTime)

	shoppingGroup := r.Group("/shopping")

	{
		shoppingGroup.GET("/index", ShopIdxHandle)
		shoppingGroup.GET("/home", ShopHomeHandle)
	}
}
