package session

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func ExeCookie() {
	r := gin.Default()

	r.GET("/cookie", func(c *gin.Context) {
		cookie, err := c.Cookie("key_cookie")

		if err != nil {
			cookie = "NotSet"
			c.SetCookie("key_cookie", "value_cookie", 60, "/", "localhost", false, true)
		}
		fmt.Printf("cookie value is:%s\n", cookie)
	})

	if err := r.Run(":20000"); err != nil {
		fmt.Println("run programmer is wrong!")
	}
}
