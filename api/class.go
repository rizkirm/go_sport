package api

import (
	"database/sql"
	"net/http"
	db "sports/db/sqlc"

	"github.com/gin-gonic/gin"
)

type createClassRequest struct {
	ClientID    int32  `json:"client_id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (server *Server) createClass(ctx *gin.Context) {
	var req createClassRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	arg := db.CreateClassParams{
		ClientID:    req.ClientID,
		Code:        req.Code,
		Name:        req.Name,
		Description: req.Description,
	}
	class, err := server.store.CreateClass(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, class)
}

type getClassRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getClass(ctx *gin.Context) {
	var req getClassRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	class, err := server.store.GetClass(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, class)
}

type listClassRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listClass(ctx *gin.Context) {
	var req listClassRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListClassParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	class, err := server.store.ListClass(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, class)
}
