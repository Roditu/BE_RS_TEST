package api

import (
	db "github.com/Roditu/BE_RS_TEST/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/person/add", server.createPerson)
	router.GET("/person/:id", server.getPerson)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorRespone(err error) gin.H {
	return gin.H{"error": err.Error()}
}