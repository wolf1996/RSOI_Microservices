package main

import (
	"github.com/wolf1996/gateway/queuemanager/manager"
	"github.com/spf13/viper"
	"github.com/wolf1996/gateway/resources/eventsclient"
	"github.com/wolf1996/gateway/resources/userclient"
)

func parseViper() manager.Config {
	viper.ReadInConfig()
	user := viper.GetString("rabbit.user")
	pass := viper.GetString("rabbit.password")
	addres := viper.GetString("rabbit.addres")
	eventServiceAddres := viper.GetString("event_info_service.addres")
	userServiceAddres := viper.GetString("user_info_service.addres")
	return manager.Config{manager.RabbitConfig{addres,user,pass}, eventsclient.Config{eventServiceAddres,eventsclient.QConfig{addres,user,pass}},
		userclient.Config{userServiceAddres,userclient.QConfig{addres,user,pass}}}
}

func prepareViper()  {
	viper.SetConfigName("worker_config")
	viper.AddConfigPath("/etc/gateway/")
	viper.AddConfigPath("/home/ksg/disk_d/labs2017M/RSOI/2/src/github.com/wolf1996/gateway/")
}

func main() {
	prepareViper()
	conf := parseViper()
	manager.StartApplication(conf)
	select{}
}
