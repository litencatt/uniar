package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Top(c *gin.Context) {
	c.HTML(http.StatusOK, "top/index.go.tmpl", gin.H{
		"title": "Main Index",
	})
}
