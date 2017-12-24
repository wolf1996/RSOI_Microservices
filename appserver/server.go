package appserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wolf1996/gateway/appserver/controllers"
	"github.com/wolf1996/gateway/resources/eventsclient"
	"github.com/wolf1996/gateway/resources/registrationclient"
	"github.com/wolf1996/gateway/resources/userclient"
	"github.com/wolf1996/gateway/resources/authclient"
	"github.com/wolf1996/gateway/appserver/middleware"
	"github.com/wolf1996/gateway/token"
)

type GatewayConfig struct {
	Port                 string
	UserInfoConf         userclient.Config
	EventsInfoConf       eventsclient.Config
	RegistrationInfoConf registrationclient.Config
	AuthConf             authclient.Config
}

var port string

func applyConfig(config GatewayConfig) {
	port = config.Port
	userclient.SetConfigs(config.UserInfoConf)
	eventsclient.SetConfigs(config.EventsInfoConf)
	registrationclient.SetConfigs(config.RegistrationInfoConf)
	authclient.SetConfigs(config.AuthConf)
}

func StartServer(config GatewayConfig) error {
	applyConfig(config)
	router := gin.Default()
	auth := router.Group("/", middleware.TokenAuth())

	auth.GET("/hello", func(c *gin.Context) {
		tkn := c.MustGet(middleware.AtokenName).(token.Token)
		user := tkn.LogIn
		var respMsg string
		respMsg = "hello " + user
		c.JSON(http.StatusOK, gin.H{"message": respMsg})
	})
	auth.POST("/events/:event_id/register", controllers.RegistrateMe)
	auth.GET("/user/registrations/:pagenum", controllers.GetRegistrations)
	auth.GET("/user/registrations/", controllers.GetRegistrations)
	auth.DELETE("/registrations/:registration_id/remove", controllers.RemoveRegistration)
	auth.GET("/registrations/:registration_id", controllers.GetRegisrationInfo)	
	auth.GET("/user/info", controllers.GetUserInfo)

	router.GET("/event/:event_id", controllers.GetEventInfo)
	router.GET("/events/:pagenum", controllers.GetEvents)
	router.POST("/login",controllers.LogIn)
	router.Run(port)
	return nil
}
