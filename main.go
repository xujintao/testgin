package main

import (
	"flag"

	"github.com/gin-gonic/gin"
	"github.com/xujintao/glog"
	"github.com/xujintao/testgin/routers"
)

func init() {
	flag.Set("log_dir", "./log")
	flag.Set("v", "3")
	flag.Set("alsologtostderr", "true")
	flag.Parse()
}

func main() {
	//创建日志文件
	gin.DefaultWriter = glog.CopyStandardLogTo("INFO")

	// gin.SetMode(gin.ReleaseMode)

	r := routers.SetupRouter()

	r.Run(":8080")
	glog.Flush()
}
