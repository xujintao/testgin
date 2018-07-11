package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/xujintao/testgin/config"
	"github.com/xujintao/testgin/etcd3"
	"github.com/xujintao/testgin/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/resolver"
)

func Testgrpc(ctx *gin.Context) {
	r := etcd3.NewResolver(fmt.Sprintf("%s:%d", config.ETCDIp, config.ETCDPort))
	resolver.Register(r)

	conn, err := grpc.Dial(r.Scheme()+"://author/hello_service", grpc.WithInsecure(), grpc.WithBalancerName(roundrobin.Name))
	if err != nil {
		log.Panic(err)
	}

	client := pb.NewHelloServiceClient(conn)

	// 测试负载均衡
	// ticker := time.NewTicker(2 * time.Second)
	// for t := range ticker.C {
	// 	resp, err := client.SayHello(context.Background(), &pb.HelloRequest{Greeting: "world"})
	// 	if err == nil {
	// 		log.Printf("%v: Reply is %s\n", t, resp.Reply)
	// 	} else {
	// 		log.Println(err)
	// 	}
	// }

	//响应http
	cancelCtx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp, err := client.SayHello(cancelCtx, &pb.HelloRequest{Greeting: "world"})
	if err != nil {
		log.Panic(err)
	}
	log.Printf("Reply is %s\n", resp.Reply)
	ctx.String(http.StatusOK, resp.Reply)
}
