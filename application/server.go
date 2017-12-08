package application

import (
	"github.com/gin-gonic/gin"
	"github.com/wolf1996/frontend/application/controller"
	"github.com/wolf1996/frontend/application/client"
	"github.com/gin-contrib/multitemplate"
	"path/filepath"
	"fmt"
)

type Config struct {
	Port      string
	Static    string
	Backend   string
	Templates string
}

func validateConfig(config Config)(err error) {
	return
}

func createRenderer(templatesDir string)(multitemplate.Render){
	r := multitemplate.New()

	layouts, err := filepath.Glob( "templates/layouts/*.html")
	if err != nil {
		panic(err.Error())
	}

	includes, err := filepath.Glob( "templates/includes/*.html")
	if err != nil {
		panic(err.Error())
	}

	// Generate our templates2 map from our layouts/ and includes/ directories
	for _, layout := range layouts {

		files := append([]string{"templates/includes/base.html"}, includes...)
		files = append(files, layout)
		fmt.Print(files)
		fmt.Print(filepath.Base(layout))
		r.AddFromFiles(filepath.Base(layout), files...)
	}
	return r
}

func StartApplication(config Config)(err error){
	err = validateConfig(config)
	if err != nil {
		return
	}
	client.ApplyConfig(client.Config{config.Backend})
	router := gin.Default()
	router.HTMLRender = createRenderer(config.Templates)
	router.Static("/static",config.Static)
	router.GET("/events/:page_num",controller.EventsList)
	router.GET("/event/:event",controller.GetEventInfo)
	router.GET("/events/",controller.EventsList)
	router.POST("/events/register",controller.RegistreMe)
	router.POST("/regs/remove",controller.RemoveReg)
	router.GET("/user_info",controller.GetUserInfo)
	router.GET("/user_regs",controller.GetUserRegistrations)
	router.GET("/user_regs/:page_num",controller.GetUserRegistrations)
	err = router.Run(config.Port)
	return
}
