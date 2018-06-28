package config

import (
	"flag"
	"log"
	"os"

	"github.com/jinzhu/configor"
)

var BConfig = &Config{}

var (
	DBName     string
	DBUser     string
	DBPassword string
	DBIp       string
	DBPort     uint
	DBTable    string
)

type Config struct {
	DB struct {
		Name     string
		User     string `default:"root"`
		Password string `required:"true" default:"root"`
		Ip       string
		Port     uint `default:"3306"`
		Table    string
	}
}

func init() {
	//配置文件检查
	if len(os.Args) == 1 {
		log.Fatal("没有指定配置文件，程序退出")
	}

	filePath := os.Args[1]
	if !fileExists(filePath) {
		log.Fatal("指定的配置文件不存在，程序退出")
	}

	//解析配置文件
	configor.Load(BConfig, filePath)
	DBName = BConfig.DB.Name
	DBUser = BConfig.DB.User
	DBPassword = BConfig.DB.Password
	DBIp = BConfig.DB.Ip
	DBPort = BConfig.DB.Port
	DBTable = BConfig.DB.Table

	//日志配置
	flag.Set("log_dir", "./log")
	flag.Set("v", "3")
	flag.Set("alsologtostderr", "true")
	flag.Parse()
}

func fileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
