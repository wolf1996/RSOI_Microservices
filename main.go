package main

import (
	"github.com/spf13/viper"
	"github.com/wolf1996/frontend/application"
	"log"
)

func parseViper() application.Config {
	viper.ReadInConfig()
	port := viper.GetString("port")
	backend := viper.GetString("backend")
	templates := viper.GetString("templates")
	static := viper.GetString("static")
	return application.Config{port, static, backend, templates}
}

func prepareViper()  {
	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/frontend/")
	viper.AddConfigPath("/home/ksg/disk_d/labs2017M/RSOI/2/src/github.com/wolf1996/frontend/")
	viper.SetDefault("port", ":8080")
}

func main() {
	prepareViper()
	conf := parseViper()
	log.Fatal(application.StartApplication(conf).Error())
}
