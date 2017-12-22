package main

import (
	"github.com/spf13/viper"
	"github.com/wolf1996/gateway/appserver"
	"github.com/wolf1996/gateway/resources/eventsclient"
	"github.com/wolf1996/gateway/resources/registrationclient"
	"github.com/wolf1996/gateway/resources/userclient"
)

func parseViper() appserver.GatewayConfig {
	viper.ReadInConfig()
	port := viper.GetString("port")
	userServiceAddres := viper.GetString("user_info_service.addres")
	eventServiceAddres := viper.GetString("event_info_service.addres")
	registrationServiceAddres := viper.GetString("registration_info_service.addres")
	user := viper.GetString("event_info_service.events_rabbit.user")
	pass := viper.GetString("event_info_service.events_rabbit.password")
	addres := viper.GetString("event_info_service.events_rabbit.addres")
	eventQue := eventsclient.QConfig{addres, user, pass}
	user = viper.GetString("user_info_service.users_rabbit.user")
	pass = viper.GetString("user_info_service.users_rabbit.password")
	addres = viper.GetString("user_info_service.users_rabbit.addres")
	crtUser := viper.GetString("user_info_service.crt")
	crtEvent := viper.GetString("event_info_service.crt")
	crtRegs := viper.GetString("registration_info_service.crt")
	userQue := userclient.QConfig{addres, user, pass}
	return appserver.GatewayConfig{port, userclient.Config{userServiceAddres, crtUser, userQue},
		eventsclient.Config{eventServiceAddres, crtEvent, eventQue}, registrationclient.Config{registrationServiceAddres, crtRegs}}
}

func prepareViper() {
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
