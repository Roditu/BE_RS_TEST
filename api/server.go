package api

import (
	db "github.com/Roditu/BE_RS_TEST/db/sqlc"
	"github.com/Roditu/BE_RS_TEST/util"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store *db.Store
	router *gin.Engine
	tokenMaker *util.JWTMaker
}

func NewServer(store *db.Store, maker *util.JWTMaker) *Server {
	server := &Server{store: store, tokenMaker: maker}
	router := gin.Default()

	router.POST("/person/add", server.createPerson)
	router.GET("/person/:id", server.getPerson)

	router.POST("/register", server.CreateUser)
	router.POST("/login", server.Login)

	auth := router.Group("/").Use(server.AuthMiddleware())
    {
        auth.POST("/tasks", server.AddTask)
        auth.GET("/tasks", server.ListTasks)
        auth.POST("/tasks/:id/finish", server.FinishTask)
    }

	user := router.Group("/").Use(server.AuthMiddleware())
    {
        user.GET("/user", server.getUserByUserId)
    }
	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorRespone(err error) gin.H {
	return gin.H{"error": err.Error()}
}