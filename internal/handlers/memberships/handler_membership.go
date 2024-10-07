package memberships

import (
	"context"

	"github.com/MuhFrds/simple-forum-go/internal/model/memberships"
	"github.com/gin-gonic/gin"
)

type membershipService interface{
	SignUp(ctx context.Context, req memberships.SignUpRequest) error
}

type Handler struct {
	*gin.Engine
 
	membershipSvc membershipService
}

func NewHandler(api *gin.Engine,  membershipSvc membershipService) *Handler{
	return	&Handler{
		Engine: api,
		membershipSvc: membershipSvc,
	}
} 

func (h *Handler) RegisterRoute(){
	route:= h.Group("memberships")
	route.GET("/ping",h.Ping)
	route.POST("/signup", h.SignUp)
}