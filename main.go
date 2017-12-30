package main

import (
	"github.com/spf13/viper"
	"github.com/wolf1996/auth/application"
	"github.com/wolf1996/auth/application/models"
	"github.com/wolf1996/auth/application/storage"
	"github.com/wolf1996/auth/application/tokenanager"
	"github.com/wolf1996/stats/client"
)

func parseViper() application.Config {
	viper.ReadInConfig()
	port := viper.GetString("port")
	salt := viper.GetString("database.salt")
	database_user := viper.GetString("database.username")
	pass := viper.GetString("database.password")
	dbname := viper.GetString("database.dbname")
	addres := viper.GetString("database.addres")
	storageaddres := viper.GetString("storage.addres")
	tokensalt := viper.GetString("token.salt")
	accesstokenexptime := viper.GetInt64("token.access_token_exparision")
	refreshtokenexptime := viper.GetInt64("token.refresh_token_exparision")
	codeflowexptime := viper.GetInt64("token.codeflow_exparision")
	var statsconf client.Config
	statsconf.ProducerName = viper.GetString("stats.name")
	statsconf.Sconfig.BufferSize = viper.GetInt("stats.handler.buffer_size")
	statsconf.Sconfig.Rabbit.Addres = viper.GetString("stats.handler.rabbit.addres")
	statsconf.Sconfig.Rabbit.User = viper.GetString("stats.handler.rabbit.user")
	statsconf.Sconfig.Rabbit.Pass = viper.GetString("stats.handler.rabbit.password")
	statsconf.ChConfig.BufferSize  = viper.GetInt("stats.response.buffer_size")
	statsconf.ChConfig.Rabbit.Addres = viper.GetString("stats.response.rabbit.addres")
	statsconf.ChConfig.Rabbit.User = viper.GetString("stats.response.rabbit.user")
	statsconf.ChConfig.Rabbit.Pass = viper.GetString("stats.response.rabbit.password")
	statsconf.TimeOut = viper.GetInt64("stats.timeout")
	statsconf.RedConf.HashName = viper.GetString("stats.redis.hash_name")
	statsconf.RedConf.Addres = viper.GetString("stats.redis.addres")
	statsconf.RedConf.Db = viper.GetInt("stats.redis.db")

	crt := viper.GetString("crt")
	key := viper.GetString("key")
	return application.Config{port, crt, key, storage.Config{storageaddres},models.DatabaseConfig{
		database_user, pass, dbname, addres,salt}, tokenanager.Config{
		tokensalt, accesstokenexptime, refreshtokenexptime, codeflowexptime,
	}, statsconf}
}

func prepareViper() {
	viper.SetConfigName("config")
	viper.AddConfigPath("/etc/auth/")
	viper.AddConfigPath("/home/ksg/disk_d/labs2017M/RSOI/2/src/github.com/wolf1996/auth/")
	viper.SetDefault("port", ":8003")
}

func main() {
	prepareViper()
	conf := parseViper()
	application.StartApplication(conf)
}