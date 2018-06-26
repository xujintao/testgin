package main

import (
	"github.com/gin-gonic/gin"
	"github.com/xujintao/glog"
	"github.com/xujintao/testgin/routers"
)

func main() {
	//创建日志文件
	gin.DefaultWriter = glog.CopyStandardLogTo("INFO")

	// gin.SetMode(gin.ReleaseMode)

	r := routers.SetupRouter()

	r.Run(":8080")
	glog.Flush()
}
