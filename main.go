package main

import (
	"github.com/wolf1996/gateway/appserver"
	"github.com/spf13/viper"
	"github.com/wolf1996/gateway/appserver/resources"
)

func parseViper() appserver.GatewayConfig {
	viper.ReadInConfig()
	port := viper.GetString("port")
	userServiceAddres := viper.GetString("user_info_service.addres")
	return appserver.GatewayConfig{port, resources.UserInfoConfig{userServiceAddres}}
}

func prepareViper()  {
	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/gateway/")
	viper.AddConfigPath("/home/ksg/disk_d/labs2017M/RSOI/2/src/github.com/wolf1996/gateway/")
	viper.SetDefault("port", ":8080")
}

func main() {
	prepareViper()
	conf := parseViper()
	appserver.StartServer(conf)
}