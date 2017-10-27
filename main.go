package main

import (
	"github.com/wolf1996/user/application"
	"github.com/spf13/viper"
)




func parseViper() application.UserConfig {
	viper.ReadInConfig()
	port := viper.GetString("port")
	return application.UserConfig{port}
}

func prepareViper()  {
	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/gateway/")
	viper.AddConfigPath("/home/ksg/disk_d/labs2017M/RSOI/2/src/github.com/wolf1996/user/")
	viper.SetDefault("port", ":8000")
}

func main()  {
	prepareViper()
	conf := parseViper()
	application.StartApplication(conf)
}
