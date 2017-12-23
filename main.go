package main

import (
	"github.com/spf13/viper"
	"github.com/wolf1996/auth/application"
	"github.com/wolf1996/auth/application/models"
	"github.com/wolf1996/auth/application/storage"
	"github.com/wolf1996/auth/application/tokenanager"
)

func parseViper() application.Config {
	viper.ReadInConfig()
	port := viper.GetString("port")
	salt := viper.GetString("salt")
	database_user := viper.GetString("database.username")
	pass := viper.GetString("database.password")
	dbname := viper.GetString("database.dbname")
	addres := viper.GetString("database.addres")
	storageaddres := viper.GetString("storage.addres")
	tokensalt := viper.GetString("token.salt")
	accesstokenexptime := viper.GetInt64("token.access_token_exparision")
	refreshtokenexptime := viper.GetInt64("token.refresh_token_exparision")
	crt := viper.GetString("crt")
	key := viper.GetString("key")
	return application.Config{port, crt, key, storage.Config{storageaddres},models.DatabaseConfig{
		database_user, pass, dbname, addres,salt}, tokenanager.Config{
		tokensalt, accesstokenexptime, refreshtokenexptime,
	}}
}

func prepareViper() {
	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/gateway/")
	viper.AddConfigPath("/home/ksg/disk_d/labs2017M/RSOI/2/src/github.com/wolf1996/events/")
	viper.SetDefault("port", ":8001")
}

func main() {
	prepareViper()
	conf := parseViper()
	application.StartApplication(conf)
}