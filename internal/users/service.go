package users

import (
	"context"
	"time"

	"github.com/naldeco98/YourTicket/internal/domain"
)

type Service interface {
	/*
		Register: creates a new user in the database

		Condition: username must be unique and password must be hashed

		Description:
		Look for the user in the database
		If the user exists, return an error
		If the user does not exist, create a new user
	*/
	Register(ctx context.Context, user *domain.User) (int, error)
}
type service struct {
	r Repository
}

// NewService creates a new user service
// It requires a user repository 'r'
func NewService(r Repository) Service {
	return &service{r: r}
}

func (s *service) Register(ctx context.Context, user *domain.User) (int, error) {
	var (
		err error
		id  int
	)
	id, err = s.r.LookByUsername(user.Username)
	if id != 0 {
		return 0, err
	}
	user.CreatedAt = time.Now()
	id, err = s.r.Create(user)
	if err != nil {
		return 0, err
	}
	return id, nil
}
