package user_repo

import (
	"context"
	"errors"
	"go-api/internal/model"
	"go-api/internal/repository"
)

var (
	userDB = []model.User{
		{ID: "1", Name: "mai1", Age: 21},
		{ID: "2", Name: "mai2", Age: 22},
		{ID: "3", Name: "mai3", Age: 23},
		{ID: "4", Name: "mai4", Age: 24},
		{ID: "5", Name: "mai5", Age: 25},
		{ID: "6", Name: "mai6", Age: 26},
		{ID: "7", Name: "mai7", Age: 27},
		{ID: "8", Name: "mai8", Age: 28},
		{ID: "9", Name: "mai9", Age: 29},
		{ID: "10", Name: "mai10", Age: 30},
	}
)

type userRepository struct{}

func NewUserRepository() repository.UserRepository {
	return &userRepository{}
}

func (r *userRepository) Save(ctx context.Context, user *model.User) error {
	userDB = append(userDB, *user)
	return nil
}

func (r *userRepository) Update(ctx context.Context, userUpdate *model.User, id string) error {
	for i, user := range userDB {
		if user.ID == id {
			userDB[i] = user
			return nil
		}
	}
	return errors.New("update fail")
}

func (r *userRepository) Delete(ctx context.Context, id string) error {
	for i, user := range userDB {
		if user.ID == id {
			userDB = append(userDB[:i], userDB[i+1:]...)
			return nil
		}
	}
	return errors.New("delete fail")
}

func (r *userRepository) Find(ctx context.Context) ([]model.User, error) {
	return userDB, nil
}

func (r *userRepository) FindOne(ctx context.Context, id string) (*model.User, error) {
	for _, user := range userDB {
		if user.ID == id {
			return &user, nil
		}
	}
	return nil, errors.New("user not found")
}
