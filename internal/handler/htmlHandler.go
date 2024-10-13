package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Serve the dashboard
func ServeRealTime(c *gin.Context) {
	c.HTML(http.StatusOK, "realtime.html", nil)
}
