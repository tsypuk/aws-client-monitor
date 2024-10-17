package router

import (
	"aws-client-monitor/docs"
	"aws-client-monitor/internal/handler"
	"github.com/gin-contrib/cors"
	//"github.com/gin-gonic/contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func CreateRouter(router *gin.Engine) *gin.Engine {
	docs.SwaggerInfo.BasePath = "/api/v1"

	// CORS middleware setup
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		AllowCredentials: true,
	}))
	v1 := router.Group("/api/v1")
	{
		v1.GET("/status", handler.StatusHandler)
		//eg := v1.Group("/example")
		//{
		//	eg.GET("/helloworld", Helloworld)
		//}
	}

	router.GET("/ws", handler.WsHandler)
	router.GET("/realtime", handler.ServeRealTime)
	router.GET("/dashboard", handler.ServeDashBoard)
	router.GET("/dashboard-layout.html", handler.ServeDashBoardLayout)

	//router.LoadHTMLFiles("templates/realtime.html", "templates/dashboard.html", "templates/dashboard-layout.html")
	// Load HTML templates
	//router.LoadHTMLGlob("templates/*")
	//router.LoadHTMLFiles("templates/dashboard.html")
	//router.LoadHTMLFiles("templates/dashboard-layout.html")

	//router.Static("/static", "./static")
	router.Use(static.Serve("/", static.LocalFile("./react-admin/build", true)))

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return router
}
