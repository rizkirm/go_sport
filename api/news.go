package api

import (
	"database/sql"
	"net/http"
	db "sports/db/sqlc"

	"github.com/gin-gonic/gin"
)

type createNewsRequest struct {
	Code        string `json:"code"`
	Title       string `json:"title"`
	Description string `json:"description"`
	PathPhoto   string `json:"path_photo"`
}

func (server *Server) createNews(ctx *gin.Context) {
	var req createNewsRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateNewsParams{
		Code:        req.Code,
		Title:       req.Title,
		Description: req.Description,
		PathPhoto:   req.PathPhoto,
	}

	news, err := server.store.CreateNews(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, news)
}

type getNewsRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getNews(ctx *gin.Context) {
	var req getNewsRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	news, err := server.store.GetNews(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, news)
}

type listNewsRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listNews(ctx *gin.Context) {
	var req listNewsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListNewsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	news, err := server.store.ListNews(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, news)
}
