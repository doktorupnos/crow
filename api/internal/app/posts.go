package app

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/doktorupnos/crow/api/internal/database"
	"github.com/doktorupnos/crow/api/internal/respond"
	"github.com/google/uuid"
)

type CreatePostRequest struct {
	Body string `json:"body"`
}

func (s *State) CreatePost(w http.ResponseWriter, r *http.Request, user database.User) error {
	var req CreatePostRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		return APIError{
			Code: http.StatusBadRequest,
			Err:  err,
		}
	}

	// TODO: validate request

	now := time.Now()
	post, err := s.db.CreatePost(r.Context(), database.CreatePostParams{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		Body:      req.Body,
		UserID:    user.ID,
	})
	if err != nil {
		return err
	}

	log.Printf("%#v\n", post)

	w.WriteHeader(http.StatusCreated)

	return nil
}

type Pagination struct {
	Limit  int32
	Offset int32
}

func extractPagination(r *http.Request) (Pagination, error) {
	queries := r.URL.Query()

	limit := 10
	if queries.Has("limit") {
		str := queries.Get("limit")
		if str != "null" {
			l, err := strconv.Atoi(str)
			if err != nil {
				return Pagination{}, err
			}
			limit = l
		}
	}

	page := 0
	if queries.Has("page") {
		p, err := strconv.Atoi(queries.Get("page"))
		if err != nil {
			return Pagination{}, err
		}
		page = p
	}

	return Pagination{
		Limit:  int32(limit),
		Offset: int32(page) * int32(limit),
	}, nil
}

type PostResponse struct {
	ID          uuid.UUID `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Body        string    `json:"body"`
	UserID      uuid.UUID `json:"user_id"`
	UserName    string    `json:"user_name"`
	Likes       int64     `json:"likes"`
	LikedByUser bool      `json:"liked_by_user"`
	Self        bool      `json:"self"`
}

func (s *State) GetPosts(w http.ResponseWriter, r *http.Request, user database.User) error {
	if r.URL.Query().Has("u") {
		return s.getPostsByUser(w, r, user)
	}
	return s.getAllPosts(w, r, user)
}

func (s *State) getAllPosts(w http.ResponseWriter, r *http.Request, user database.User) error {
	pages, err := extractPagination(r)
	if err != nil {
		return APIError{
			Code: http.StatusBadRequest,
			Err:  err,
		}
	}

	posts, err := s.db.GetPosts(r.Context(), database.GetPostsParams{
		Limit:  pages.Limit,
		Offset: pages.Offset,
	})
	if err != nil {
		return err
	}

	response := make([]PostResponse, 0, len(posts))
	for _, post := range posts {
		likes, err := s.db.GetLikesForPost(r.Context(), post.ID)
		if err != nil {
			return err
		}

		likedByUser, err := s.db.UserLikesPost(r.Context(), database.UserLikesPostParams{
			UserID: user.ID,
			PostID: post.ID,
		})
		if err != nil {
			return err
		}

		response = append(response, PostResponse{
			ID:          post.ID,
			CreatedAt:   post.CreatedAt,
			UpdatedAt:   post.UpdatedAt,
			Body:        post.Body,
			UserID:      post.UserID,
			UserName:    post.UserName,
			Likes:       likes,
			LikedByUser: likedByUser,
			Self:        post.UserID == user.ID,
		})
	}

	respond.JSON(w, http.StatusOK, response)

	return nil
}

func (s *State) getPostsByUser(w http.ResponseWriter, r *http.Request, user database.User) error {
	pages, err := extractPagination(r)
	if err != nil {
		return APIError{
			Code: http.StatusBadRequest,
			Err:  err,
		}
	}

	posts, err := s.db.GetPostsByUser(r.Context(), database.GetPostsByUserParams{
		Name:   r.URL.Query().Get("u"),
		Limit:  pages.Limit,
		Offset: pages.Offset,
	})
	if err != nil {
		return err
	}

	response := make([]PostResponse, 0, len(posts))
	for _, post := range posts {
		likes, err := s.db.GetLikesForPost(r.Context(), post.ID)
		if err != nil {
			return err
		}

		likedByUser, err := s.db.UserLikesPost(r.Context(), database.UserLikesPostParams{
			UserID: user.ID,
			PostID: post.ID,
		})
		if err != nil {
			return err
		}

		response = append(response, PostResponse{
			ID:          post.ID,
			CreatedAt:   post.CreatedAt,
			UpdatedAt:   post.UpdatedAt,
			Body:        post.Body,
			UserID:      post.UserID,
			UserName:    post.UserName,
			Likes:       likes,
			LikedByUser: likedByUser,
			Self:        post.UserID == user.ID,
		})
	}

	respond.JSON(w, http.StatusOK, response)

	return nil
}

func (s *State) DeletePost(w http.ResponseWriter, r *http.Request, user database.User) error {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		return APIError{
			Code: http.StatusBadRequest,
			Err:  err,
		}
	}

	err = s.db.DeletePost(r.Context(), database.DeletePostParams{
		ID:     id,
		UserID: user.ID,
	})
	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusOK)
	return nil
}
