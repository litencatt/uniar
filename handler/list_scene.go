package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListScene(c *gin.Context) {
	c.HTML(http.StatusOK, "scenes/index.go.tmpl", gin.H{
		"title": "Scenes Index",
	})
}
