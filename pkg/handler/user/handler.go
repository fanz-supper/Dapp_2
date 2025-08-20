package userhandler

import (
	"Dapp_2/pkg/handler"
	"Dapp_2/pkg/service/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	handler.BaseHandler
	service *userservice.Service
}

func (h *Handler) registerGroup() *gin.RouterGroup {
	return h.Server.Gin.Group(h.Group)
}

func (h *Handler) routes() http.Handler {
	//add user resultful api

	return h.Server.Gin
}

func NewHandler(server *handler.Server, group string, userService *userservice.Service) *Handler {

	uh := &Handler{
		BaseHandler: handler.BaseHandler{
			Server: server,
			Group:  group,
			Router: server.Gin.Group(group),
		},
		service: userService,
	}

	uh.routes()

	return uh
}
