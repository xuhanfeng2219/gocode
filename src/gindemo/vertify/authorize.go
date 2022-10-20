package vertify

import (
	"fmt"
	"github.com/casbin/casbin"
	xormadapter "github.com/casbin/xorm-adapter"
	"github.com/gin-gonic/gin"
)

func main() {
	adapter := xormadapter.NewAdapter("mysql", "root@root@tcp(127.0.0.1:3306)/goblog?charset=utf8", true)

	e := casbin.NewEnforcer("conf/rbac_models.conf", adapter)

	if err := e.LoadPolicy(); err != nil {
		fmt.Println("load policy failed!")
	}

	r := gin.New()

	r.POST("/api/v1/add", func(c *gin.Context) {
		fmt.Println("add policy")
		e.AddPolicy("admin", "/api/v1/hello", "GET")
	})

	r.DELETE("/api/v1/delete", func(c *gin.Context) {
		fmt.Println("delete policy")
		e.RemovePolicy("admin", "/api/v1/hello", "GET")
	})

	r.GET("/api/v1/get", func(c *gin.Context) {
		fmt.Println("cat policy")
		list := e.GetPolicy()
		for _, l := range list {
			for _, v := range l {
				fmt.Printf("value: %s,", v)
			}
		}
	})

	r.Use(Authorize(e))

	r.GET("/api/v1/hello", func(c *gin.Context) {
		fmt.Println("hello receieve get request!")
	})

	if err := r.Run(":90000"); err != nil {
		fmt.Println("run wrong!")
	}
}

func Authorize(e *casbin.Enforcer) gin.HandlerFunc {
	return func(c *gin.Context) {
		obj := c.Request.URL.RequestURI()
		act := c.Request.Method
		sub := "admin"

		e.Enforce(sub, obj, act)
	}
}
