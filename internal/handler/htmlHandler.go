package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Serve the dashboard
func ServeDashboard(c *gin.Context) {
	c.HTML(http.StatusOK, "dashboard.html", nil)
}
