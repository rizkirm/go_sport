package api

import (
	"database/sql"
	"net/http"
	db "sports/db/sqlc"

	"github.com/gin-gonic/gin"
)

type createClientRequest struct {
	Code             string `json:"code"`
	Name             string `json:"name"`
	DomainLocal      string `json:"domain_local"`
	DomainProduction string `json:"domain_production"`
	WhatsappNumber   string `json:"whatsapp_number"`
	WhatsappMessage  string `json:"whatsapp_message"`
	WhatsappLink     string `json:"whatsapp_link"`
	FbLink           string `json:"fb_link"`
	IgLink           string `json:"ig_link"`
	PathLogo         string `json:"path_logo"`
}

func (server *Server) createClient(ctx *gin.Context) {
	var req createClientRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateClientParams{
		Code:             req.Code,
		Name:             req.Name,
		DomainLocal:      req.DomainLocal,
		DomainProduction: req.DomainProduction,
		WhatsappNumber:   req.WhatsappNumber,
		WhatsappMessage:  req.WhatsappMessage,
		WhatsappLink:     req.WhatsappLink,
		FbLink:           req.FbLink,
		IgLink:           req.IgLink,
		PathLogo:         req.PathLogo,
	}

	client, err := server.store.CreateClient(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, client)
}

type getClientRequest struct {
	ID int32 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getClient(ctx *gin.Context) {
	var req getClientRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	client, err := server.store.GetClient(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, client)
}

type listClientRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listClient(ctx *gin.Context) {
	var req listClientRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListClientsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	client, err := server.store.ListClients(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, client)
}

type updateClientRequest struct {
	Name             string `json:"name"`
	DomainLocal      string `json:"domain_local"`
	DomainProduction string `json:"domain_production"`
	WhatsappNumber   string `json:"whatsapp_number"`
	WhatsappMessage  string `json:"whatsapp_message"`
	WhatsappLink     string `json:"whatsapp_link"`
	FbLink           string `json:"fb_link"`
	IgLink           string `json:"ig_link"`
	PathLogo         string `json:"path_logo"`
}

func (server *Server) updateClient(ctx *gin.Context) {
	var req updateClientRequest
	var updateClientReq getClientRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := ctx.ShouldBindUri(&updateClientReq); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateClientParams{
		ID:               updateClientReq.ID,
		Name:             req.Name,
		DomainLocal:      req.DomainLocal,
		DomainProduction: req.DomainProduction,
		WhatsappNumber:   req.WhatsappNumber,
		WhatsappMessage:  req.WhatsappMessage,
		WhatsappLink:     req.WhatsappLink,
		FbLink:           req.FbLink,
		IgLink:           req.IgLink,
		PathLogo:         req.PathLogo,
	}

	client, err := server.store.UpdateClient(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, client)
}

func (server *Server) deleteClient(ctx *gin.Context) {
	var req getClientRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.store.DeleteClient(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, "Succes delete a client")
}
