package main

import (
	"fmt"
	"log"
	"net/http"
)

// @Description  go语言学习练习
// @Author playclouds
// @Update    2021/8/1 17:11

func serverApp1() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Hello Go")
	})
	if err := http.ListenAndServe("0.0.0.0:8080", mux); err != nil {
		log.Fatal(err)
	}
}

func serverDeBug1() {
	if err := http.ListenAndServe("0.0.0.0:8081", http.DefaultServeMux); err != nil {
		log.Fatal(err)
	}
}

func main() {
	go serverDeBug1()
	go serverApp1()
	select {}

}
