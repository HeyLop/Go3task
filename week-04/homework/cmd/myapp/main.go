package main

import "fmt"
import "github.com/gin-gonic/gin"

// @Description  go3task-week04
// @Author playclouds
// @Update    2021/8/8 20:02

func main() {
	fmt.Println("这里是main函数！")
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello go3task - week04",
		})
	})
	r.Run()
}
