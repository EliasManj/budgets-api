package api

import (
	db "github.com/eliasmanj/budgets-api/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  db.Querier
	Router *gin.Engine
}

func NewServer(store db.Querier) (*Server, error) {
	router := gin.Default()
	server := &Server{
		store:  store,
		Router: router,
	}

	// Routes
	router.POST("/users", server.CreateUser)
	router.GET("/users/login", server.LoginUser)
	router.DELETE("/users/:username")

	return server, nil

}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
