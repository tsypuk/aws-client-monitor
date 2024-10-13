package router

import (
	"aws-client-monitor/docs"
	"aws-client-monitor/internal/handler"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func CreateRouter(router *gin.Engine) *gin.Engine {
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := router.Group("/api/v1")
	{
		v1.GET("/status", handler.StatusHandler)
		//eg := v1.Group("/example")
		//{
		//	eg.GET("/helloworld", Helloworld)
		//}
	}

	router.GET("/ws", handler.WsHandler)
	router.GET("/", handler.ServeRealTime)
	router.GET("/dashboard", handler.ServeDashBoard)

	router.LoadHTMLFiles("templates/realtime.html")
	router.LoadHTMLFiles("templates/dashboard.html")

	router.Static("/static", "./static")

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return router
}
