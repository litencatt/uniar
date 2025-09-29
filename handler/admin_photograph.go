package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/litencatt/uniar/repository"
	"github.com/litencatt/uniar/service"
)

type AdminPhotographHandler struct {
	PhotographService *service.Photgraph
	SearchService     *service.SearchService
}

func (h *AdminPhotographHandler) ListPhotograph(c *gin.Context) {
	ctx := c.Request.Context()

	// ドロップダウン選択肢用のデータを取得
	groups, err := h.SearchService.Querier.GetGroup(ctx, h.SearchService.DB)
	if err != nil {
		groups = []repository.GetGroupRow{} // エラー時は空配列
	}

	photoTypes, err := h.SearchService.Querier.GetPhotoTypeList(ctx, h.SearchService.DB)
	if err != nil {
		photoTypes = []string{} // エラー時は空配列
	}

	// 検索パラメータを取得
	var searchParams service.PhotographSearchParams
	if err := c.ShouldBindQuery(&searchParams); err == nil {
		// 検索パラメータがある場合は検索実行
		if searchParams.Name != "" || searchParams.GroupID != 0 || searchParams.PhotoType != "" {
			photographs, err := h.SearchService.SearchPhotograph(ctx, searchParams)
			if err != nil {
				c.HTML(http.StatusInternalServerError, "500.go.tmpl", gin.H{
					"error": err.Error(),
				})
				return
			}

			c.HTML(http.StatusOK, "admin/photograph_list.go.tmpl", gin.H{
				"photographs":   photographs,
				"searchParams":  searchParams,
				"groups":        groups,
				"photoTypes":    photoTypes,
			})
			return
		}
	}

	// 検索パラメータがない場合は全件取得
	photographs, err := h.PhotographService.ListAllForAdmin(ctx)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "500.go.tmpl", gin.H{
			"error": err.Error(),
		})
		return
	}

	c.HTML(http.StatusOK, "admin/photograph_list.go.tmpl", gin.H{
		"photographs":  photographs,
		"searchParams": service.PhotographSearchParams{},
		"groups":       groups,
		"photoTypes":   photoTypes,
	})
}


func (h *AdminPhotographHandler) EditPhotograph(c *gin.Context) {
	ctx := c.Request.Context()

	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.HTML(http.StatusBadRequest, "400.go.tmpl", gin.H{
			"error": "Invalid photograph ID",
		})
		return
	}

	photograph, err := h.PhotographService.GetByID(ctx, id)
	if err != nil {
		c.HTML(http.StatusNotFound, "404.go.tmpl", gin.H{
			"error": "Photograph not found",
		})
		return
	}

	c.HTML(http.StatusOK, "admin/photograph_edit.go.tmpl", gin.H{
		"photograph": photograph,
	})
}

func (h *AdminPhotographHandler) UpdatePhotograph(c *gin.Context) {
	ctx := c.Request.Context()

	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.HTML(http.StatusBadRequest, "400.go.tmpl", gin.H{
			"error": "Invalid photograph ID",
		})
		return
	}

	name := c.PostForm("name")
	groupID, _ := strconv.ParseInt(c.PostForm("group_id"), 10, 64)
	photoType := c.PostForm("photo_type")
	abbreviation := c.PostForm("abbreviation")
	nameForOrder := c.PostForm("name_for_order")

	params := service.UpdatePhotographParams{
		Name:         name,
		GroupID:      groupID,
		PhotoType:    photoType,
		Abbreviation: abbreviation,
		NameForOrder: nameForOrder,
	}

	err = h.PhotographService.Update(ctx, id, params)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "500.go.tmpl", gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Redirect(http.StatusFound, "/admin/photograph")
}

func (h *AdminPhotographHandler) DeletePhotograph(c *gin.Context) {
	ctx := c.Request.Context()

	idParam := c.Param("id")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		c.HTML(http.StatusBadRequest, "400.go.tmpl", gin.H{
			"error": "Invalid photograph ID",
		})
		return
	}

	err = h.PhotographService.Delete(ctx, id)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "500.go.tmpl", gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Redirect(http.StatusFound, "/admin/photograph")
}

func (h *AdminPhotographHandler) NewPhotograph(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/photograph_new.go.tmpl", gin.H{})
}

func (h *AdminPhotographHandler) AddPhotograph(c *gin.Context) {
	ctx := c.Request.Context()

	name := c.PostForm("name")
	groupID, _ := strconv.ParseInt(c.PostForm("group_id"), 10, 64)
	photoType := c.PostForm("photo_type")

	params := service.AddPhotographParams{
		Name:      name,
		GroupID:   groupID,
		PhotoType: photoType,
	}

	err := h.PhotographService.Add(ctx, params)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "500.go.tmpl", gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Redirect(http.StatusFound, "/admin/photograph")
}