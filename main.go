package main

import (
	"github.com/wolf1996/gateway/appserver"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")
	viper.AddConfigPath("")
	appserver.StartServer()
}