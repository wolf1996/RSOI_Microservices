package main

import (
	"github.com/wolf1996/user/application"
	"github.com/spf13/viper"
	"github.com/wolf1996/user/application/models"
)




func parseViper() application.Config {
	viper.ReadInConfig()
	port := viper.GetString("port")
	database_user := viper.GetString("database.username")
	pass := viper.GetString("database.password")
	dbname := viper.GetString("database.dbname")
	addres := viper.GetString("database.addres")
	crt :=  viper.GetString("crt")
	key :=  viper.GetString("key")
	return application.Config{port, crt,key,models.DatabaseConfig{
		database_user,pass, dbname, addres,
	}}
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
