package module

import (
	"context"
	"go-api/internal/model"
)

type UserService interface {
	Create(ctx context.Context, user *model.User) error
	Update(ctx context.Context, userUpdate *model.User, id string) error
	Delete(ctx context.Context, id string) error

	GetAll(ctx context.Context) ([]model.User, error)
	GetOne(ctx context.Context, id string) (*model.User, error)
}

