package main

// @Description  go3task-week05
// @Author playclouds
// @Update    2021/8/15 22:24

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

var MaxQueue map[string][]int64
var ok bool

func LimitFreqSingle(qn string, count uint, timeLockWindow int64) bool {
	currTime := time.Now().Unix()
	if MaxQueue == nil {
		MaxQueue = make(map[string][]int64)
	}
	if _, ok = MaxQueue[qn]; !ok {
		MaxQueue[qn] = make([]int64, 0)
	}

	if uint(len(MaxQueue[qn])) < count {
		MaxQueue[qn] = append(MaxQueue[qn], currTime)
		return true
	}
	earlyTime := MaxQueue[qn][0]
	if currTime-earlyTime <= timeLockWindow {
		return false
	} else {
		MaxQueue[qn] = MaxQueue[qn][1:]
		MaxQueue[qn] = append(MaxQueue[qn], currTime)
	}
	return true
}
func main() {

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		ip := c.ClientIP()
		qn := fmt.Sprintf("client ip:%s", ip)
		if !LimitFreqSingle(qn, 10, 10) {
			c.JSON(200, gin.H{
				"code": 400,
				"msg":  fmt.Sprintf("Current ip:%s  access too often", ip),
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "hello go3task - week05",
		})
	})
	r.Run()

}
