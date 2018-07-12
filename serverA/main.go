package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

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
	if err := etcd3.Register(etcdAddr, serviceName, serverAddr, 5); err != nil {
		log.Fatal(err)
	}

	//优雅关闭
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT)
	go func() {
		s := <-ch
		log.Printf("receive signal '%v'", s)
		if err := etcd3.UnRegister(serviceName, serverAddr); err != nil {
			log.Panic(err)
		}

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

	if err := s.Serve(l); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

type helloServer struct{}

func (helloServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("getting request from client.\n")
	return &pb.HelloResponse{Reply: "Hello, " + req.Greeting}, nil
}
