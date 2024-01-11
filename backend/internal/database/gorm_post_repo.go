package database

import (
	"github.com/doktorupnos/crow/backend/internal/like"
	"github.com/doktorupnos/crow/backend/internal/pages"
	"github.com/doktorupnos/crow/backend/internal/post"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GormPostRepo struct {
	db *gorm.DB
}

func NewGormPostRepo(db *gorm.DB) *GormPostRepo {
	return &GormPostRepo{db}
}

func (r *GormPostRepo) Create(p post.Post) error {
	return r.db.Create(&p).Error
}

func (r *GormPostRepo) Load(params post.LoadParams) ([]post.FeedPost, error) {
	var posts []post.Post
	err := r.db.
		Order("created_at " + params.Order).
		Scopes(pages.Paginate(params.PaginationParams)).
		Find(&posts).Error
	if err != nil {
		return nil, err
	}

	feedPosts := make([]post.FeedPost, 0, len(posts))
	for _, p := range posts {
		var username string

		err := r.db.
			Table("users").
			Select("name").
			Where("id = ?", p.UserID).
			Scan(&username).
			Error
		if err != nil {
			return nil, err
		}

		var likes int64
		if err := r.db.Model(&like.PostLike{}).
			Where("post_id = ?", p.ID).
			Count(&likes).
			Error; err != nil {
			return nil, err
		}

		var c int64
		if err := r.db.Model(&like.PostLike{}).
			Where("post_id = ? AND user_id = ?", p.ID, params.UserID).
			Count(&c).
			Error; err != nil {
			return nil, err
		}

		var likedByUser bool
		if c == 1 {
			likedByUser = true
		}

		feedPosts = append(feedPosts, post.FeedPost{
			Post:        p,
			UserName:    username,
			Likes:       likes,
			LikedByUser: likedByUser,
		})
	}

	return feedPosts, err
}

func (r *GormPostRepo) LoadAllByID(params post.LoadParams) ([]post.FeedPost, error) {
	var posts []post.Post
	err := r.db.
		Order("created_at DESC").
		Scopes(pages.Paginate(params.PaginationParams)).
		Where("user_id = ?", params.UserID).
		Find(&posts).Error
	if err != nil {
		return nil, err
	}

	feedPosts := make([]post.FeedPost, 0, len(posts))
	for _, p := range posts {
		var username string

		err := r.db.
			Table("users").
			Select("name").
			Where("id = ?", p.UserID).
			Scan(&username).
			Error
		if err != nil {
			return nil, err
		}

		var likes int64
		if err := r.db.Model(&like.PostLike{}).
			Where("post_id = ?", p.ID).
			Count(&likes).
			Error; err != nil {
			return nil, err
		}

		var c int64
		if err := r.db.Model(&like.PostLike{}).
			Where("post_id = ? AND user_id = ?", p.ID, params.UserID).
			Count(&c).
			Error; err != nil {
			return nil, err
		}

		var likedByUser bool
		if c == 1 {
			likedByUser = true
		}

		feedPosts = append(feedPosts, post.FeedPost{
			Post:        p,
			UserName:    username,
			Likes:       likes,
			LikedByUser: likedByUser,
		})
	}

	return feedPosts, err
}

func (r *GormPostRepo) LoadByID(id uuid.UUID) (post.Post, error) {
	var p post.Post
	err := r.db.First(&p, id).Error
	return p, err
}

func (r *GormPostRepo) Update(p post.Post) error {
	return r.db.Save(&p).Error
}

func (r *GormPostRepo) Delete(p post.Post) error {
	return r.db.Delete(&p).Error
}
