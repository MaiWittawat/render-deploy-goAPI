package repository

import (
	"context"
	"go-api/internal/model"
)

type UserRepository interface {
	Save(ctx context.Context, user *model.User) error
	Update(ctx context.Context, userUpdate *model.User, id string) error
	Delete(ctx context.Context, id string) error

	Find(ctx context.Context) ([]model.User, error)
	FindOne(ctx context.Context, id string) (*model.User, error)
}
