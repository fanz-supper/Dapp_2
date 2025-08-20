package handler

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	Gin *gin.Engine
}

func NewServer() *Server {
	ginEngine := gin.New()
	ginEngine.Use(gin.Logger())
	ginEngine.Use(gin.Recovery())

	return &Server{
		Gin: ginEngine,
	}
}

func (s *Server) Run() error {
	return s.Gin.Run(":8080")
}
