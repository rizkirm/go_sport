package api

import (
	"database/sql"
	"net/http"
	db "sports/db/sqlc"

	"github.com/gin-gonic/gin"
)

type createClientCmsRequest struct {
	HeroSection  string `json:"hero_section"`
	AboutSection string `json:"about_section"`
}

func (server *Server) createClientCMS(ctx *gin.Context) {
	var req createClientCmsRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateClientCMSParams{
		HeroSection:  req.HeroSection,
		AboutSection: req.AboutSection,
	}

	clientCMS, err := server.store.CreateClientCMS(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, clientCMS)
}

type getClientCMSRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getClientCMS(ctx *gin.Context) {
	var req getClientCMSRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	clientCMS, err := server.store.GetClientCMS(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, clientCMS)
}

type listClientCMSRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listClientCMS(ctx *gin.Context) {
	var req listClientCMSRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListClientCMSParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	clientCMS, err := server.store.ListClientCMS(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, clientCMS)
}
