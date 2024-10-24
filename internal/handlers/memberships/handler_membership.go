package memberships

import (
	"context"

	"github.com/MuhFrds/simple-forum-go/internal/middleware"
	"github.com/MuhFrds/simple-forum-go/internal/model/memberships"
	"github.com/gin-gonic/gin"
)

type membershipService interface{
	SignUp(ctx context.Context, req memberships.SignUpRequest) error
	Login(ctx context.Context, req memberships.LoginRequest) (string, string, error)
	ValidateRefreshToken(ctx context.Context, userID int64, request memberships.RefreshTokenRequest) (string, error)
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
	route.POST("/login", h.Login)

	routeRefresh := h.Group("memberships")
	routeRefresh.Use(middleware.AuthMiddleware())
	routeRefresh.POST("/refresh", h.Refresh)
}