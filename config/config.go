package config

import (
	"flag"

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
	//解析配置文件
	configor.Load(BConfig, "config/config.json")
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
