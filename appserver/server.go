package appserver

import ("github.com/gin-gonic/gin"
	"net/http"
	"github.com/wolf1996/gateway/appserver/controllers"
	"github.com/wolf1996/gateway/appserver/resources"
)

type GatewayConfig struct{
	Port string
	UserInfoConf resources.UserInfoConfig
}

var port string

func applyConfig(config GatewayConfig){
	port = config.Port
	resources.SetUserInfoConfigs(config.UserInfoConf)
}

func StartServer(config GatewayConfig) error  {
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
		respMsg = "hello "+ user
		c.JSON(http.StatusOK, gin.H{"message":respMsg})
	})

	auth.GET("/user_info", controllers.GetUserInfo)

	router.Run(port)
	return nil
}
