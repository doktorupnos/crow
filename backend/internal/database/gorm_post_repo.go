package database

import (
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

func (r *GormPostRepo) Load(params post.PaginationParams) ([]post.FeedPost, error) {
	var posts []post.Post
	err := r.db.Scopes(post.Paginate(params)).Find(&posts).Error
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

		feedPosts = append(feedPosts, post.FeedPost{
			Post:     p,
			UserName: username,
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
