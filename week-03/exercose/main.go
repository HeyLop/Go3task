package main

import (
	"fmt"
	"net/http"
	"runtime"
)

// @Description  go语言学习练习
// @Author playclouds
// @Update    2021/8/1 21:08

func main() {

	go func() {
		mux := http.NewServeMux()
		mux.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
			query := req.URL.Query()
			singlePanic := query.Get("num")
			fmt.Println("DeBug Page singlePanic Num:", singlePanic)
			fmt.Fprintln(resp, fmt.Sprintf("DeBug Page singlePanic Num: %s", singlePanic))
			if singlePanic == "1" {
				//panic("Panic by manually")
				//os.Exit(1)
				fmt.Println("DeBug Page singlePanic Num:", singlePanic, "Panic by manually")
				//模拟协程退出，但是好不行.....

				runtime.Goexit()

			}

		})

		http.ListenAndServe("0.0.0.0:8081", mux)

	}()

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Hello Go")
		fmt.Println("Hello Go Page")
	})
	http.ListenAndServe("0.0.0.0:8080", mux)

}
