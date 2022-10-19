package data

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
	"net/http"
)

//xml/json/xml/yaml/类似于java的properties和protobuf

func DataHandler() {
	r := gin.Default()
	//json
	r.GET("/json", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": "json", "status": http.StatusOK})
	})
	//struct
	r.GET("/struct", func(c *gin.Context) {
		var msg struct {
			Name   string
			Msg    string
			Number int
		}

		msg.Name = "root"
		msg.Msg = "message"
		msg.Number = 123

		c.JSON(http.StatusOK, msg)
	})

	//xml
	r.GET("/xml", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{"message": "abc"})
	})

	//yaml
	r.GET("/yaml", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{"name": "zhangsan"})
	})

	//protobuf
	r.GET("/protobuf", func(c *gin.Context) {
		resp := []int64{int64(1), int64(2)}
		label := "label"

		data := &protoexample.Test{
			Label: &label,
			Reps:  resp,
		}
		c.ProtoBuf(http.StatusOK, data)
	})

	if err := r.Run(":8000"); err != nil {
		fmt.Println("hello is wrong, please check some errors out!")
	}
}
