package user

import (
	"context"
	"strings"
	"time"

	"github.com/doktorupnos/crow/api/internal/app/passwd"
	"github.com/doktorupnos/crow/api/internal/database"
	"github.com/google/uuid"
)

type Service interface {
	Create(context.Context, CreateRequest) (database.User, error)
}

type PostgresService struct {
	db *database.Queries
}

func NewPostgresService(db *database.Queries) *PostgresService {
	return &PostgresService{db: db}
}

var _ Service = (*PostgresService)(nil)

func (r *PostgresService) Create(ctx context.Context, req CreateRequest) (database.User, error) {
	if err := req.Validate(); err != nil {
		return database.User{}, err
	}

	hashed, err := passwd.Hash(req.Password)
	if err != nil {
		return database.User{}, err
	}

	now := time.Now()
	user, err := r.db.CreateUser(ctx, database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		Name:      req.Name,
		Password:  hashed,
	})
	if err != nil {
		if strings.Contains(err.Error(), "violates unique constraint") {
			return database.User{}, ErrNameTaken
		}
		return database.User{}, err
	}

	return user, nil
}
