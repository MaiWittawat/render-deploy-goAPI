package user_svc

import (
	"context"
	"go-api/internal/model"
	"go-api/internal/module"
	"go-api/internal/repository"
)


type userService struct {
	repo repository.UserRepository
}


func NewUserService(repo repository.UserRepository) module.UserService {
	return &userService{
		repo: repo,
	}
}

func (s *userService) Create(ctx context.Context, user *model.User) error {
	return s.repo.Save(ctx, user)
}

func (s *userService) Update(ctx context.Context, userUpdate *model.User, id string) error {
	return s.repo.Update(ctx, userUpdate, id)
}

func (s *userService) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}

func (s *userService) GetAll(ctx context.Context) ([]model.User, error) {
	users, err := s.repo.Find(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil	
}

func (s *userService) GetOne(ctx context.Context, id string) (*model.User, error) {
	user, err := s.repo.FindOne(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}