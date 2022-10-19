package vertify

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Person struct {
	Age      int       `form:"age" binding:"required,gt=10"`
	Name     string    `form:"name" binding:"required"`
	Birthday time.Time `form:"Birthday" time_format:"2022-10-08" time_utc:"1"`
}

func main() {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		var person Person
		if err := c.ShouldBind(&person); err != nil {
			c.String(http.StatusInternalServerError, fmt.Sprint(err))
			return
		}
		c.String(http.StatusOK, fmt.Sprintf("%#v", person))
	})

	if err := r.Run(":2000"); err != nil {
		fmt.Println("sprint name wrong!")
	}
}
