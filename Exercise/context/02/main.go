package main

import (
	"context"
	"fmt"
	"time"
)

// @Description  go语言学习练习
// @Author playclouds
// @Update    2021/8/1 14:46

const shortDuration = 1 * time.Millisecond

func main() {
	d := time.Now().Add(shortDuration)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		fmt.Println("ctx.Err()", ctx.Err())
	}
}
