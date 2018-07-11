package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/xujintao/glog"
	"github.com/xujintao/testgin/config"
	"github.com/xujintao/testgin/etcd3"
	"github.com/xujintao/testgin/pb"
	"google.golang.org/grpc"
)

var (
	serviceName = "hello_service"
)

func main() {
	//日志写到glog
	glog.CopyStandardLogTo("INFO")

	//注册服务
	etcdAddr := fmt.Sprintf("%s:%d", config.ETCDIp, config.ETCDPort)
	serverAddr := fmt.Sprintf("%s:%d", config.ServerAIp, config.ServerAPort)
	go etcd3.Register(etcdAddr, serviceName, serverAddr, 5)

	//优雅关闭
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT)
	go func() {
		s := <-ch
		log.Printf("receive signal '%v'", s)
		etcd3.UnRegister(serviceName, serverAddr)

		if i, ok := s.(syscall.Signal); ok {
			os.Exit(int(i))
		} else {
			os.Exit(0)
		}
	}()

	//启动服务
	port := config.ServerAPort
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("starting hello service at %d", port)
	s := grpc.NewServer()
	defer s.GracefulStop()
	pb.RegisterHelloServiceServer(s, &helloServer{})

	if err := s.Serve(tcpKeepAliveListener{l.(*net.TCPListener)}); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

// tcpKeepAliveListener sets TCP keep-alive timeouts on accepted
// connections. It's used by ListenAndServe and ListenAndServeTLS so
// dead TCP connections (e.g. closing laptop mid-download) eventually
// go away.
type tcpKeepAliveListener struct {
	*net.TCPListener
}

func (ln tcpKeepAliveListener) Accept() (c net.Conn, err error) {
	tc, err := ln.AcceptTCP()
	if err != nil {
		return
	}
	tc.SetKeepAlive(true)
	tc.SetKeepAlivePeriod(3 * time.Minute)
	return tc, nil
}

type helloServer struct{}

func (helloServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("getting request from client.\n")
	return &pb.HelloResponse{Reply: "Hello, " + req.Greeting}, nil
}
