package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func HomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World",
	})
}

func PostTest(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Test Post",
	})
}

func QueryString(c *gin.Context) {
	query := c.Query("query")
	c.JSON(200, gin.H{
		"message": query,
	})
}

func ParamsTest(c *gin.Context) {
	name := c.Param("name")
	age := c.Param("age")
	c.JSON(200, gin.H{
		"name": name,
		"age":  age,
	})
}

func main() {
	fmt.Println("Hello seikai")
	r := gin.Default()
	r.GET("/", HomePage)
	r.POST("/", PostTest)
	r.GET("/query", QueryString)
	r.GET("/path/:name/:age", ParamsTest)
	r.Run()
}
