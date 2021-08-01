package main

import (
	"fmt"
	"net/http"
)

// @Description  go语言学习练习
// @Author playclouds
// @Update    2021/8/1 17:06

func serverApp() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Hello Go")
	})
	http.ListenAndServe("0.0.0.0:8080", mux)
}

func serverDeBug() {
	http.ListenAndServe("0.0.0.0:8081", http.DefaultServeMux)
}

func main() {
	go serverDeBug()
	serverApp()

}
