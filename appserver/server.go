package appserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wolf1996/gateway/appserver/controllers"
	"github.com/wolf1996/gateway/resources/eventsclient"
	"github.com/wolf1996/gateway/resources/registrationclient"
	"github.com/wolf1996/gateway/resources/userclient"
)

type GatewayConfig struct {
	Port                 string
	UserInfoConf         userclient.Config
	EventsInfoConf       eventsclient.Config
	RegistrationInfoConf registrationclient.Config
}

var port string

func applyConfig(config GatewayConfig) {
	port = config.Port
	userclient.SetConfigs(config.UserInfoConf)
	eventsclient.SetConfigs(config.EventsInfoConf)
	registrationclient.SetConfigs(config.RegistrationInfoConf)
}

func StartServer(config GatewayConfig) error {
	applyConfig(config)
	router := gin.Default()
	auth := router.Group("/", gin.BasicAuth(
		gin.Accounts{
			"simpleUser": "1",
			"eventOwner": "1",
		}))

	auth.GET("/hello", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)
		var respMsg string
		respMsg = "hello " + user
		c.JSON(http.StatusOK, gin.H{"message": respMsg})
	})
	auth.POST("/events/:event_id/register", controllers.RegistrateMe)
	auth.GET("/user/registrations/:pagenum", controllers.GetRegistrations)
	auth.GET("/user/registrations/", controllers.GetRegistrations)
	auth.POST("/registrations/:registration_id/remove", controllers.RemoveRegistration)

	auth.GET("/user/info", controllers.GetUserInfo)
	router.GET("/events/:event_id", controllers.GetEventInfo)
	router.GET("/registrations/:registration_id", controllers.GetRegisrationInfo)

	router.Run(port)
	return nil
}
