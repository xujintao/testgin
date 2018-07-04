package controllers

import (
	"net/http"
	"time"
)

var httpClient *http.Client

func init() {
	tp := http.DefaultTransport.(*http.Transport)
	// tp.TLSClientConfig = &tls.Config{InsecureSkipVerify: true} //不验证证书

	httpClient = &http.Client{
		// 默认连接池配置
		Transport: tp,

		// 会话超时时间
		Timeout: 10 * time.Second,
	}
}
