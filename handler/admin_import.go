package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/litencatt/uniar/service"
)

type AdminImportHandler struct {
	ImportService *service.ImportService
}

func (h *AdminImportHandler) ImportCSVForm(c *gin.Context) {
	entityType := c.Param("entity")
	if entityType == "" {
		c.HTML(http.StatusBadRequest, "error/400.go.tmpl", gin.H{
			"title":   "Bad Request",
			"message": "Entity type is required",
		})
		return
	}

	validEntities := map[string]bool{
		"music":      true,
		"photograph": true,
		"scene":      true,
	}

	if !validEntities[entityType] {
		c.HTML(http.StatusBadRequest, "error/400.go.tmpl", gin.H{
			"title":   "Bad Request",
			"message": "Invalid entity type",
		})
		return
	}

	c.HTML(http.StatusOK, "admin/import_csv.go.tmpl", gin.H{
		"title":      "CSV Import",
		"entityType": entityType,
	})
}

func (h *AdminImportHandler) ImportCSVUpload(c *gin.Context) {
	entityType := c.Param("entity")
	if entityType == "" {
		c.HTML(http.StatusBadRequest, "error/400.go.tmpl", gin.H{
			"title":   "Bad Request",
			"message": "Entity type is required",
		})
		return
	}

	file, err := c.FormFile("csv_file")
	if err != nil {
		c.HTML(http.StatusBadRequest, "error/400.go.tmpl", gin.H{
			"title":   "Bad Request",
			"message": "CSV file is required",
		})
		return
	}

	src, err := file.Open()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error/500.go.tmpl", gin.H{
			"title":   "Internal Server Error",
			"message": "Failed to open CSV file",
		})
		return
	}
	defer src.Close()

	previewModeStr := c.PostForm("preview_mode")
	previewMode, _ := strconv.ParseBool(previewModeStr)

	var result *service.ImportResult

	switch entityType {
	case "music":
		result, err = h.ImportService.ImportMusicFromCSV(c.Request.Context(), src, previewMode)
	case "photograph":
		result, err = h.ImportService.ImportPhotographFromCSV(c.Request.Context(), src, previewMode)
	case "scene":
		result, err = h.ImportService.ImportSceneFromCSV(c.Request.Context(), src, previewMode)
	default:
		c.HTML(http.StatusBadRequest, "error/400.go.tmpl", gin.H{
			"title":   "Bad Request",
			"message": "Invalid entity type",
		})
		return
	}

	if err != nil {
		c.HTML(http.StatusInternalServerError, "error/500.go.tmpl", gin.H{
			"title":   "Internal Server Error",
			"message": err.Error(),
		})
		return
	}

	if previewMode {
		c.HTML(http.StatusOK, "admin/import_preview.go.tmpl", gin.H{
			"title":       "Import Preview",
			"entityType":  entityType,
			"result":      result,
			"previewMode": true,
		})
		return
	}

	if result.Failed > 0 {
		c.HTML(http.StatusOK, "admin/import_result.go.tmpl", gin.H{
			"title":      "Import Result",
			"entityType": entityType,
			"result":     result,
			"hasErrors":  true,
		})
		return
	}

	c.HTML(http.StatusOK, "admin/import_result.go.tmpl", gin.H{
		"title":      "Import Result",
		"entityType": entityType,
		"result":     result,
		"hasErrors":  false,
	})
}

func (h *AdminImportHandler) ImportCSVConfirm(c *gin.Context) {
	entityType := c.Param("entity")
	if entityType == "" {
		c.HTML(http.StatusBadRequest, "error/400.go.tmpl", gin.H{
			"title":   "Bad Request",
			"message": "Entity type is required",
		})
		return
	}

	file, err := c.FormFile("csv_file")
	if err != nil {
		c.HTML(http.StatusBadRequest, "error/400.go.tmpl", gin.H{
			"title":   "Bad Request",
			"message": "CSV file is required",
		})
		return
	}

	src, err := file.Open()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "error/500.go.tmpl", gin.H{
			"title":   "Internal Server Error",
			"message": "Failed to open CSV file",
		})
		return
	}
	defer src.Close()

	var result *service.ImportResult

	switch entityType {
	case "music":
		result, err = h.ImportService.ImportMusicFromCSV(c.Request.Context(), src, false)
	case "photograph":
		result, err = h.ImportService.ImportPhotographFromCSV(c.Request.Context(), src, false)
	case "scene":
		result, err = h.ImportService.ImportSceneFromCSV(c.Request.Context(), src, false)
	default:
		c.HTML(http.StatusBadRequest, "error/400.go.tmpl", gin.H{
			"title":   "Bad Request",
			"message": "Invalid entity type",
		})
		return
	}

	if err != nil {
		c.HTML(http.StatusInternalServerError, "error/500.go.tmpl", gin.H{
			"title":   "Internal Server Error",
			"message": err.Error(),
		})
		return
	}

	c.HTML(http.StatusOK, "admin/import_result.go.tmpl", gin.H{
		"title":      "Import Result",
		"entityType": entityType,
		"result":     result,
		"hasErrors":  result.Failed > 0,
	})
}