package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"syscall"
)

// @Description  go3task week-03
// @Author playclouds
// @Update    2021/8/1 19:40

type httpServer struct {
}

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

func regSignal(stop chan struct{}) error {
	signalChan := make(chan os.Signal, 1)
	for {
		sig := <-signalChan
		switch sig {
		case syscall.SIGHUP, syscall.SIGINT, syscall.SIGKILL:
			fmt.Println("Received exit signal, waiting for exit...")

			stop <- struct{}{}
		}
	}
}

func main() {
	stop := make(chan struct{})
	eg, _ := errgroup.WithContext(context.Background())
	eg.Go(func() error {
		var serverApp httpServer
		return server("0.0.0.0:8080", serverApp, stop)
	})
	eg.Go(func() error {
		return regSignal(stop)
	})

	if err := eg.Wait(); err != nil {
		log.Fatal(err)
	}
}
