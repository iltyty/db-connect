package conf

import (
	"fmt"
	"github.com/iltyty/db_connect/backend_go/utils"
	"github.com/spf13/viper"
	"os"
)

const configPath = "conf/config.yaml"

type Server struct {
	Host string
	Port int
}

type Database struct {
	Name     string
	Host     string
	Port     int
	Username string
	Password string
}

type Log struct {
	Dirname  string
	Filename string
}

var (
	ServerConf   = &Server{}
	DatabaseConf = &Database{}
	LogConf      = &Log{}
)

func Init() {
	viper.AddConfigPath(".")
	viper.SetConfigFile(configPath)

	dir, _ := os.Getwd()
	fmt.Println(dir)

	viper.WatchConfig()

	err := viper.ReadInConfig()
	utils.CheckError(err)

	err = viper.UnmarshalKey("server", ServerConf)
	utils.CheckError(err)
	err = viper.UnmarshalKey("database", DatabaseConf)
	utils.CheckError(err)
	err = viper.UnmarshalKey("log", LogConf)
	utils.CheckError(err)
}
