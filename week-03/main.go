package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
)

// @Description  go语言学习练习
// @Author playclouds
// @Update    2021/8/1 19:40

func main() {
	eg, _ := errgroup.WithContext(context.Background())
	eg.Go(func() error {
		mux := http.NewServeMux()
		mux.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
			query := req.URL.Query()
			singlePanic := query.Get("num")
			fmt.Println("DeBug Page singlePanic Num:", singlePanic)
			fmt.Fprintln(resp, fmt.Sprintf("DeBug Page singlePanic Num: %s", singlePanic))
			if singlePanic == "1" {
				//panic("Panic by manually")
				log.Fatal("Panic by manually")
			}

		})

		err := http.ListenAndServe("0.0.0.0:8081", mux)
		return err
	})
	eg.Go(func() error {
		mux := http.NewServeMux()
		mux.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
			fmt.Fprintln(resp, "Hello Go")
			fmt.Println("Hello Go Page")
		})
		err := http.ListenAndServe("0.0.0.0:8080", mux)
		return err
	})

	if err := eg.Wait(); err != nil {
		log.Fatal(err)
	}
}
