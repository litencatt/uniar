package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/litencatt/uniar/service"
)

type AdminMusicHandler struct {
	MusicService *service.Music
}

func (h *AdminMusicHandler) ListMusic(c *gin.Context) {
	ctx := c.Request.Context()

	musics, err := h.MusicService.ListAll(ctx)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "500.go.tmpl", gin.H{
			"error": err.Error(),
		})
		return
	}

	c.HTML(http.StatusOK, "admin/music_list.go.tmpl", gin.H{
		"musics": musics,
	})
}


func (h *AdminMusicHandler) EditMusic(c *gin.Context) {
	ctx := c.Request.Context()

	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.HTML(http.StatusBadRequest, "400.go.tmpl", gin.H{
			"error": "Invalid music ID",
		})
		return
	}

	music, err := h.MusicService.GetByID(ctx, id)
	if err != nil {
		c.HTML(http.StatusNotFound, "404.go.tmpl", gin.H{
			"error": "Music not found",
		})
		return
	}

	c.HTML(http.StatusOK, "admin/music_edit.go.tmpl", gin.H{
		"music": music,
	})
}

func (h *AdminMusicHandler) UpdateMusic(c *gin.Context) {
	ctx := c.Request.Context()

	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.HTML(http.StatusBadRequest, "400.go.tmpl", gin.H{
			"error": "Invalid music ID",
		})
		return
	}

	name := c.PostForm("name")
	normal, _ := strconv.ParseInt(c.PostForm("normal"), 10, 64)
	pro, _ := strconv.ParseInt(c.PostForm("pro"), 10, 64)
	master, _ := strconv.ParseInt(c.PostForm("master"), 10, 64)
	length, _ := strconv.ParseInt(c.PostForm("length"), 10, 64)
	colorTypeID, _ := strconv.ParseInt(c.PostForm("color_type_id"), 10, 64)
	liveID, _ := strconv.ParseInt(c.PostForm("live_id"), 10, 64)
	musicBonus, _ := strconv.ParseInt(c.PostForm("music_bonus"), 10, 64)

	params := service.UpdateMusicParams{
		Name:        name,
		Normal:      normal,
		Pro:         pro,
		Master:      master,
		Length:      length,
		ColorTypeID: colorTypeID,
		LiveID:      liveID,
		MusicBonus:  musicBonus,
	}

	err = h.MusicService.Update(ctx, id, params)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "500.go.tmpl", gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Redirect(http.StatusFound, "/admin/music")
}

func (h *AdminMusicHandler) DeleteMusic(c *gin.Context) {
	ctx := c.Request.Context()

	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.HTML(http.StatusBadRequest, "400.go.tmpl", gin.H{
			"error": "Invalid music ID",
		})
		return
	}

	err = h.MusicService.Delete(ctx, id)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "500.go.tmpl", gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Redirect(http.StatusFound, "/admin/music")
}