package handler

import (
	"github.com/gin-gonic/gin"
)

type BaseHandler struct {
	Server *Server
	Group  string
	Router *gin.RouterGroup
}
