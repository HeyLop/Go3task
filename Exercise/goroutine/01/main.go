package main

import (
	"fmt"
)

// @Description  go语言学习练习
// @Author playclouds
// @Update    2021/8/1 15:29

func echoNum() {
	for i := 0; i < 10; i++ {
		fmt.Printf("the num: %d \n", i)
	}
}

func main() {
	go echoNum()

	select {}

	//filepath.Walk()
}
