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
		//处理终止进程的信号 https://blog.csdn.net/sufwei/article/details/51610676
		//SIGHUP     终止进程     终端线路挂断
		//SIGINT     终止进程     中断进程
		//SIGKILL   终止进程     杀死进程
		//SIGTERM   终止进程     软件终止信号
		// todo 在windows下测试ctrl+c 未输出，考虑可能原因是windows下终止的信号不一样或者不是发送信号，待待在linux下测试
		case syscall.SIGHUP, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM, syscall.SIGQUIT:
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
