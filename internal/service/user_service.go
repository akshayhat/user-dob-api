package service

import (
	"context"
	"errors"
	"time"

	"github.com/akshayhat/user-dob-api/db/sqlc"
	"github.com/akshayhat/user-dob-api/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(
	ctx context.Context,
	name string,
	dob string,
) (sqlc.User, error) {

	dobTime, err := time.Parse("2006-01-02", dob)
	if err != nil {
		return sqlc.User{}, errors.New("invalid dob format, use YYYY-MM-DD")
	}

	return s.repo.CreateUser(ctx, name, dobTime)
}

func (s *UserService) GetUserByID(
	ctx context.Context,
	id int32,
) (sqlc.User, int, error) {

	user, err := s.repo.GetUserByID(ctx, id)
	if err != nil {
		return sqlc.User{}, 0, err
	}

	age := calculateAge(user.Dob)
	return user, age, nil
}

func (s *UserService) ListUsers(
	ctx context.Context,
) ([]map[string]interface{}, error) {

	users, err := s.repo.ListUsers(ctx)
	if err != nil {
		return nil, err
	}

	var result []map[string]interface{}
	for _, user := range users {
		result = append(result, map[string]interface{}{
			"id":   user.ID,
			"name": user.Name,
			"dob":  user.Dob.Format("2006-01-02"),
			"age":  calculateAge(user.Dob),
		})
	}

	return result, nil
}

func calculateAge(dob time.Time) int {
	now := time.Now()
	age := now.Year() - dob.Year()

	if now.YearDay() < dob.YearDay() {
		age--
	}
	return age
}
