package api

import (
	"database/sql"
	"net/http"
	db "sports/db/sqlc"

	"github.com/gin-gonic/gin"
)

type createClassPackageRequest struct {
	Code         string `json:"code"`
	Name         string `json:"name"`
	Price        string `json:"price"`
	CustomerType int32  `json:"customer_type"`
	Type         int32  `json:"type"`
	Description  string `json:"description"`
	PathPhoto    string `json:"path_photo"`
}

func (server *Server) createClassPackage(ctx *gin.Context) {
	var req createClassPackageRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateClassPackageParams{
		Code:         req.Code,
		Name:         req.Name,
		Price:        req.Price,
		CustomerType: req.CustomerType,
		Type:         req.Type,
		Description:  req.Description,
		PathPhoto:    req.PathPhoto,
	}

	classPackage, err := server.store.CreateClassPackage(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, classPackage)
}

type getClassPackageRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getClassPackage(ctx *gin.Context) {
	var req getClassPackageRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	classPackage, err := server.store.GetClassPackage(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, classPackage)
}

type listClassPackageRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listClassPackage(ctx *gin.Context) {
	var req listClassPackageRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListClassPackageParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	classPackage, err := server.store.ListClassPackage(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, classPackage)
}
