package api

import (
	db "sports/db/sqlc"

	"github.com/gin-gonic/gin"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/api/client", server.createClient)
	router.GET("/api/client/:id", server.getClient)
	router.GET("/api/client", server.listClient)
	router.PUT("/api/client/:id", server.updateClient)
	router.DELETE("api/client/:id", server.deleteClient)

	router.POST("/api/class", server.createClass)
	router.GET("/api/class/:id", server.getClass)
	router.GET("/api/class", server.listClass)

	router.POST("/api/news", server.createNews)
	router.GET("/api/news/:id", server.getNews)
	router.GET("/api/news", server.listNews)

	router.POST("/api/clientcms", server.createClientCMS)
	router.GET("/api/clientcms/:id", server.getClientCMS)
	router.GET("/api/clientcms", server.listClientCMS)

	router.POST("/api/classpackage", server.createClassPackage)
	router.GET("/api/classpackage/:id", server.getClassPackage)
	router.GET("/api/classpackage", server.listClassPackage)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
