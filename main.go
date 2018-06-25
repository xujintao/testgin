package main

import (
	"io"
	"log"
	"os"
	"test/testgin/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	//创建日志文件
	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file:", err)
	}
	mw := io.MultiWriter(file, os.Stdout)
	gin.DefaultWriter = mw
	gin.DefaultErrorWriter = mw
	log.SetOutput(mw)

	r := routers.SetupRouter()

	r.Run(":8080")
}
