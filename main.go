package main

import (
	"github.com/gin-gonic/gin"
	"github.com/xujintao/glog"
	"github.com/xujintao/testgin/models"
	"github.com/xujintao/testgin/routers"
)

func main() {
	//gin日志写到glog
	w := glog.CopyStandardLogTo("INFO")
	gin.DefaultWriter = w
	gin.DefaultErrorWriter = w

	defer glog.Flush()
	defer models.Close()

	// gin.SetMode(gin.ReleaseMode)

	r := routers.SetupRouter()

	r.Run(":8080")
}
