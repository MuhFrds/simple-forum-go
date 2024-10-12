package posts

import (
	"context"

	"github.com/MuhFrds/simple-forum-go/internal/middleware"
	"github.com/MuhFrds/simple-forum-go/internal/model/posts"
	"github.com/gin-gonic/gin"
)


type postService interface {
	CreatePost(ctx context.Context, UserID int64,  req posts.CreatePostRequest) error
}

type Handler struct {
	*gin.Engine
 
	postSvc postService
}

func NewHandler(api *gin.Engine,  postSvc postService) *Handler{
	return	&Handler{
		Engine: api,
		postSvc: postSvc,
	}
} 


func (h *Handler) RegisterRoute(){
	route:= h.Group("posts")
	route.Use(middleware.AuthMiddleware())

	route.POST("/create-post", h.CreatePost)
}