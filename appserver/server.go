package appserver

import ("github.com/gin-gonic/gin"
	"net/http"
)

type ServerConfig struct {
	Port string
}

func StartServer() error  {
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
	router.Run(":8080")
	return nil
}
