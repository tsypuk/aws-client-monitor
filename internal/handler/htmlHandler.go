package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Serve the dashboard
func ServeRealTime(c *gin.Context) {
	c.HTML(http.StatusOK, "realtime.html", nil)
}

func ServeDashBoard(c *gin.Context) {
	c.HTML(http.StatusOK, "dashboard.html", nil)
}

func ServeDashBoardLayout(c *gin.Context) {
	c.HTML(http.StatusOK, "dashboard-layout.html", nil)
}
