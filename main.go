package main

import (
	"github.com/spf13/viper"
	"github.com/wolf1996/gateway/appserver"
	"github.com/wolf1996/gateway/resources/eventsclient"
	"github.com/wolf1996/gateway/resources/registrationclient"
	"github.com/wolf1996/gateway/resources/userclient"
	"github.com/wolf1996/gateway/resources/authclient"
	"github.com/wolf1996/stats/client"
	"log"
	"github.com/wolf1996/gateway/resources/statsclient"
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
	authAddres := viper.GetString("auth_service.addres")
	crtAuth := viper.GetString("auth_service.crt")
	userQue := userclient.QConfig{addres, user, pass}
	var statsconf client.Config
	statsconf.ProducerName = viper.GetString("stats.name")
	statsconf.Sconfig.BufferSize = viper.GetInt("stats.handler.buffer_size")
	statsconf.Sconfig.Rabbit.Addres = viper.GetString("stats.handler.rabbit.addres")
	statsconf.Sconfig.Rabbit.User = viper.GetString("stats.handler.rabbit.user")
	statsconf.Sconfig.Rabbit.Pass = viper.GetString("stats.handler.rabbit.password")
	statsconf.Sconfig.Retries = viper.GetInt("stats.handler.retries")
	statsconf.ChConfig.BufferSize  = viper.GetInt("stats.response.buffer_size")
	statsconf.ChConfig.Rabbit.Addres = viper.GetString("stats.response.rabbit.addres")
	statsconf.ChConfig.Rabbit.User = viper.GetString("stats.response.rabbit.user")
	statsconf.ChConfig.Rabbit.Pass = viper.GetString("stats.response.rabbit.password")
	statsconf.TimeOut = viper.GetInt64("stats.timeout")
	statsconf.RedConf.HashName = viper.GetString("stats.redis.hash_name")
	statsconf.RedConf.RetriesHashName = viper.GetString("stats.redis.retries_hash_name")
	statsconf.RedConf.Addres = viper.GetString("stats.redis.addres")
	statsconf.RedConf.Db = viper.GetInt("stats.redis.db")
	var statsCliConf statsclient.Config
	statsCliConf.Addres = viper.GetString("stat_service.addres")
 	statsCliConf.Crt =  viper.GetString("stat_service.crt")
	return appserver.GatewayConfig{port, userclient.Config{userServiceAddres, crtUser, userQue},
		eventsclient.Config{eventServiceAddres, crtEvent, eventQue}, registrationclient.Config{registrationServiceAddres, crtRegs},
		authclient.Config{authAddres, crtAuth}, statsconf, statsCliConf}
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
	err := appserver.StartServer(conf)
	if err != nil {
		log.Print(err.Error())
	}
}
