package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/xujintao/testgin/config"
	"github.com/xujintao/testgin/etcd3"
	"google.golang.org/grpc/resolver"
)

var httpClient *http.Client

func init() {
	// http连接池
	tp := http.DefaultTransport.(*http.Transport)
	// tp.TLSClientConfig = &tls.Config{InsecureSkipVerify: true} //不验证证书

	httpClient = &http.Client{
		// 默认连接池配置
		Transport: tp,

		// 会话超时时间
		Timeout: 10 * time.Second,
	}

	// 劫持grpc拨号（连接池还没做）
	r := etcd3.NewResolver(fmt.Sprintf("%s:%d", config.ETCDIp, config.ETCDPort))
	resolver.Register(r)
}
