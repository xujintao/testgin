package config

import (
	"flag"
	"log"
	"os"
	"path/filepath"

	"github.com/jinzhu/configor"
)

var BConfig = &Config{}

var (
	DBName      string
	DBUser      string
	DBPassword  string
	DBIp        string
	DBPort      uint
	DBTable     string
	ETCDIp      string
	ETCDPort    uint
	ServerAIp   string
	ServerAPort uint
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

	ETCD struct {
		Ip   string
		Port uint
	}

	ServerA struct {
		Ip   string
		Port uint
	}
}

func init() {
	//配置文件检查
	if len(os.Args) == 1 {
		log.Fatal("没有指定配置文件，程序退出")
	}

	filePath := os.Args[1]
	if !fileExists(filePath) {
		log.Fatalf("指定的配置文件[%s]不存在，程序退出", filePath)
	}

	//解析配置文件
	if err := configor.Load(BConfig, filePath); err != nil {
		log.Fatalf("配置文件[%s]有问题，%s", filePath, err) //exit to fixed `sql: unknown driver "" (forgotten import?)`
	}

	DBName = BConfig.DB.Name
	DBUser = BConfig.DB.User
	DBPassword = BConfig.DB.Password
	DBIp = BConfig.DB.Ip
	DBPort = BConfig.DB.Port
	DBTable = BConfig.DB.Table
	ETCDIp = BConfig.ETCD.Ip
	ETCDPort = BConfig.ETCD.Port
	ServerAIp = BConfig.ServerA.Ip
	ServerAPort = BConfig.ServerA.Port

	//日志配置
	logPath := filepath.Join(filepath.Dir(filePath), "/log")
	os.MkdirAll(logPath, 0777)
	flag.Set("log_dir", logPath)
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
