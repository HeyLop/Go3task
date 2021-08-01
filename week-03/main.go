package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
)

// @Description  go3task week-03
// @Author playclouds
// @Update    2021/8/1 19:40

type httpServer struct {
}

type signalFunc func(<-chan struct{}) error

func (server httpServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("path is : %s \n", r.URL.Path)))
}

func server(addr string, handler http.Handler, stop <-chan struct{}) error {
	s := http.Server{
		Addr:    addr,
		Handler: handler,
	}
	go func() {
		<-stop
		s.Shutdown(context.Background())
	}()
	return s.ListenAndServe()
}

func main() {
	stop := make(chan struct{})
	eg, _ := errgroup.WithContext(context.Background())
	eg.Go(func() error {
		var serverApp httpServer
		return server("0.0.0.0:8080", serverApp, stop)
	})
	eg.Go(func() error {
		return Signal(stop)
	})

	if err := eg.Wait(); err != nil {
		log.Fatal(err)
	}
}
