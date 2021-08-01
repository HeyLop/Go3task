package main

import (
	"context"
	"fmt"
)

// @Description  go语言学习练习
// @Author playclouds
// @Update    2021/8/1 14:35

func main() {
	g := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for n := range g(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}
