package main

import (
	"github.com/spf13/viper"
	"github.com/wolf1996/stats/workerapp/worker"
	"log"
	"github.com/wolf1996/stats/application"
)

func parseViper() (worker.Config, application.Config){
	viper.ReadInConfig()
	cfg := worker.Config{}
	cfg.Hconfig.BufferSize = viper.GetInt("handler.buffer_size")
	cfg.Hconfig.NumberOfWorkers = viper.GetInt("handler.workers_number")
	cfg.Hconfig.RabConf.Addres = viper.GetString("handler.rabbit.addres")
	cfg.Hconfig.RabConf.User = viper.GetString("handler.rabbit.user")
	cfg.Hconfig.RabConf.Pass = viper.GetString("handler.rabbit.password")
	cfg.Rsconfig.BufferSize = viper.GetInt("response.buffer_size")
	cfg.Rsconfig.NumberOfWorkers = viper.GetInt("response.workers_number")
	cfg.Rsconfig.RabConf.Addres = viper.GetString("response.rabbit.addres")
	cfg.Rsconfig.RabConf.User = viper.GetString("response.rabbit.user")
	cfg.Rsconfig.RabConf.Pass = viper.GetString("response.rabbit.password")
	cfg.Dtconfig.NumberOfWorkers = viper.GetInt("databasemgr.workers_number")
	cfg.Dtconfig.BufferSize = viper.GetInt("databasemgr.buffer_size")
	cfg.Dtconfig.Types = viper.GetStringSlice("databasemgr.types")
	cfg.Dtconfig.MgoConf.Addres = viper.GetString("databasemgr.database.url")
	cfg.Dtconfig.MgoConf.DbName = viper.GetString("databasemgr.database.database")
	cfgapp := application.Config{}
	cfgapp.Crt = viper.GetString("app.crt")
	cfgapp.Key = viper.GetString("app.key")
	cfgapp.Port = viper.GetString("app.port")
	cfgapp.MgoConf.Addres = viper.GetString("app.database.url")
	cfgapp.MgoConf.DbName = viper.GetString("app.database.database")
	return cfg, cfgapp
}

func prepareViper() {
	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/gateway/")
	viper.AddConfigPath("/home/ksg/disk_d/labs2017M/RSOI/2/src/github.com/wolf1996/stats/")
	viper.SetDefault("port", ":8080")
}

func main() {
	prepareViper()
	conf, appconf := parseViper()
	err := worker.StartWorkers(conf)
	if err != nil{
		log.Fatal(err.Error())
	}
	err = application.StartApp(appconf)
	if err != nil {
		log.Fatal(err.Error())
	}
	select{}
}

