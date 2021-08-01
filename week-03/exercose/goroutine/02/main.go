package main

import (
	"fmt"
)

// @Description  go语言学习练习
// @Author playclouds
// @Update    2021/8/1 16:39

func main() {

	ch := make(chan int)

	//go func() {
	//	for i := 0; i < 10; i++ {
	//		i = <-ch
	//	}
	//	close(ch)
	//}()

	go func() {
		//time.Sleep(1*time.Second)
		val := <-ch
		fmt.Println("received a value:", val)

	}()

	//time.Sleep(10*time.Second)
	fmt.Printf("final")

}

//go run .\02\main.go --trace
