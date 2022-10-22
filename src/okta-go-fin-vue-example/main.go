package main

import (
	"github.com/gin-contrib/static"
	_ "github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var todos []string

func Lists(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"lists": todos})
}

func ListItems(c *gin.Context) {
	errMsg := "Index out of range"
	indexStr := c.Param("index")
	if index, err := strconv.Atoi(indexStr); err == nil && index < len(todos) {
		c.JSON(http.StatusOK, gin.H{"item": todos[index]})
	} else {
		if err != nil {
			errMsg = "Number expected:" + indexStr
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": errMsg})
	}
}

func AddListItem(c *gin.Context) {
	item := c.PostForm("item")
	todos = append(todos, item)
	c.String(http.StatusCreated, c.FullPath()+"/"+strconv.Itoa(len(todos)-1))
}

func main() {
	todos = append(todos, "Write the Apps")
	r := gin.Default()
	r.Use(static.Serve("/", static.LocalFile("./todo-vue/dist", false)))
	r.GET("/api/lists", Lists)
	r.GET("/api/lists/:index", ListItems)
	r.POST("/api/lists", AddListItem)
	_ = r.Run()
}
