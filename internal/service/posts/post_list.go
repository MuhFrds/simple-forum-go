package posts

import (
	"context"

	"github.com/MuhFrds/simple-forum-go/internal/model/posts"
	"github.com/rs/zerolog/log"
)

func (s *service) GetAllPost(ctx context.Context, pageSize, pageIndex int) (posts.GetAllPostResponse, error){
	limit := pageSize
	offset := pageSize * (pageIndex -1)
	response, err := s.postRepo.GetAllPost(ctx, limit, offset) 
	if err != nil {
		log.Err(err).Msg("error get all post from database") 
	}

	return response, nil
}