// routes.go

package routers

import (
	"controllers"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes() *gin.Engine {
	// Set Gin to production mode
	//gin.SetMode(gin.ReleaseMode)

	// Set the router as the default one provided by Gin
	router := gin.Default()

	//设置静态资源
	router.Static("static", "./static")
	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	router.LoadHTMLGlob("templates/*")

	//处理/demo的GET requests 
	router.GET("/demo",controllers.ShowDemo)
	// 处理ueditor的各类requests 
	router.Any("/demo/controller", controllers.Action)
