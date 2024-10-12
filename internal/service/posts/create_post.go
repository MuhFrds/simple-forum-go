package posts

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/MuhFrds/simple-forum-go/internal/model/posts"
	"github.com/rs/zerolog/log"
)

func (s *service) CreatePost(ctx context.Context, UserID int64,  req posts.CreatePostRequest) error {
	PostHashtags := strings.Join(req.PostHashtags, ",")

	now := time.Now()
	model := posts.PostModel{
		UserID : UserID,
		PostTitle : req.PostTitle,
		PostContent : req.PostContent,
		PostHashtags : PostHashtags,
		CreatedAt : now,
		UpdatedAt: now,
		CreatedBy: strconv.FormatInt(UserID, 10),
		UpdatedBy: strconv.FormatInt(UserID, 10),
	}

	err  := s.postRepo.CreatePost(ctx, model)
	if err != nil {
		log.Error().Err(err).Msg("Error Create Post to repository")
		return err
	}

	return nil
}