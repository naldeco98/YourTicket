package users

import (
	"context"
	"fmt"

	"github.com/naldeco98/YourTicket/internal/domain"
)

type Service interface {
	/*
		Create: creates a new user in the database

		Condition: username must be unique and password must be hashed

		Description:
		Look for the user in the database
		If the user exists, return an error
		If the user does not exist, create a new user
	*/
	Create(ctx context.Context, username, password string, roleId, gymId int) (int, error)
	/*
		GetById: returns a user by id

		Condition: id must be a valid integer

		Description:
		Look for the user in the database
		If the user does not exist, return an error
	*/
	GetById(ctx context.Context, id int) (domain.User, error)
	/*
		GetAll: returns all users

		Condition: none

		Description:
		Look for all users in the database
	*/
	GetAll(ctx context.Context) ([]domain.User, error)
	/*
		Update: updates a user in the database

		Condition: id must be a valid integer

		Description:
		Look for the user in the database
		If the user does not exist, return an error
		If the user exists, update the user
	*/
	Update(ctx context.Context, user domain.User) (domain.User, error)
}
type service struct {
	r Repository
}

// NewService creates a new user service
// It requires a user repository 'r'
func NewService(r Repository) Service {
	return &service{r: r}
}

func (s *service) Create(ctx context.Context, username, password string, roleId, gymId int) (int, error) {
	var (
		err error
		id  int
	)
	id, err = s.r.LookByUsername(username)
	if id != 0 {
		return 0, err
	}
	newUser := domain.User{
		Username: username,
		Password: password,
		RoleId:   roleId,
		GymId:    gymId,
	}
	id, err = s.r.Create(&newUser)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *service) GetById(ctx context.Context, id int) (domain.User, error) {
	var (
		err  error
		user domain.User
	)
	user, err = s.r.ReadByID(id)
	if err != nil {
		return user, err
	}
	return user, nil

}

func (s *service) GetAll(ctx context.Context) ([]domain.User, error) {
	var (
		err   error
		users []domain.User
	)
	users, err = s.r.ReadAll()
	if err != nil {
		return users, err
	}
	return users, nil
}

func (s *service) Update(ctx context.Context, update domain.User) (domain.User, error) {
	var (
		err error
		id  int
	)
	user, err := s.r.ReadByID(update.Id)
	if err != nil {
		return domain.User{}, err
	}
	id, _ = s.r.LookByUsername(update.Username)
	if id != 0 {
		return domain.User{}, fmt.Errorf("user with username %s already exists", update.Username)
	}
	user.Username = update.Username
	user.Password = update.Password
	user.RoleId = update.RoleId
	user.GymId = update.GymId
	if err = s.r.Update(user); err != nil {
		return domain.User{}, err
	}
	return user, nil
}
