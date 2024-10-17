package posts

import (
	"context"
	"database/sql"

	"github.com/MuhFrds/simple-forum-go/internal/model/posts"
)

func (r *repository) GetUserActivity(ctx context.Context, model posts.UserActivityModel) (*posts.UserActivityModel, error){
	query := `SELECT id, post_id, user_id,  is_liked, created_at, updated_at, created_by, updated_by
	FROM user_activities WHERE post_id = ? AND user_id = ?`

	var response posts.UserActivityModel
	row := r.db.QueryRowContext(ctx, query, model.PostID, model.UserID	)

	err:=row.Scan(&response.ID, &response.PostID, &response.UserID, &response.IsLiked, &response.CreatedAt, &response.UpdatedAt, &response.CreatedBy, &response.UpdatedBy)
	if err != nil {
		if err == sql.ErrNoRows{
			return nil, nil
		}
		return nil	, err
	}
	return &response, nil
}

func (r *repository) CreateUserActivity(ctx context.Context, model posts.UserActivityModel) error{
	query := `INSERT INTO user_activities (post_id, user_id, is_liked, created_at, updated_at, created_by, updated_by)
	 VALUES (?, ?, ?, ?, ?, ?, ?)`
	 _, err := r.db.ExecContext(ctx, query, model.PostID, model.UserID, model.IsLiked, model.CreatedAt, model.UpdatedAt, model.CreatedBy, model.UpdatedBy)
	 if err != nil {
		return nil
	 }
	 return nil
}

func (r *repository) UpdateUserActivity(ctx context.Context, model posts.UserActivityModel) error{
	query := `UPDATE user_activities SET is_liked = ?, updated_at = ?, updated_by = ? WHERE post_id = ? AND user_id = ?`
	_, err := r.db.ExecContext(ctx, query , model.IsLiked, model.UpdatedAt, model.UpdatedBy, model.PostID, model.UserID)
	if err != nil {
		return err 
	}
	return nil
}