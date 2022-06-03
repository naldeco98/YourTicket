package users

import (
	"database/sql"

	"github.com/naldeco98/YourTicket/internal/domain"
)

const (
	CREATE          = `INSERT INTO users (username, password, role_id, gym_id, created_at) VALUES (?, ?, ?, ?, ?);`
	READ            = `SELECT id, username, role_id, gym_id, created_at FROM users WHERE id = ?;`
	LOOKFORUSERNAME = `SELECT id FROM users WHERE username = ?;`
	READALL         = `SELECT id, username, role_id, gym_id, created_at FROM users;`
)

type Repository interface {
	Create(u *domain.User) (int, error)
	LookByUsername(username string) (int, error)
	ReadByID(id int) (domain.User, error)
	ReadAll() ([]domain.User, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(user *domain.User) (int, error) {
	res, err := r.db.Exec(CREATE, user.Username, user.Password, user.RoleId, user.GymId, user.CreatedAt)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (r *repository) LookByUsername(username string) (int, error) {
	var id int
	err := r.db.QueryRow(LOOKFORUSERNAME, username).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *repository) ReadByID(id int) (domain.User, error) {
	var err error
	var user domain.User
	err = r.db.QueryRow(READ, id).Scan(&user.Id, &user.Username, &user.RoleId, &user.GymId, &user.CreatedAt)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) ReadAll() ([]domain.User, error) {
	var (
		err   error
		users []domain.User
	)
	rows, err := r.db.Query(READALL)
	if err != nil {
		return users, err
	}
	for rows.Next() {
		var user domain.User
		err = rows.Scan(&user.Id, &user.Username, &user.RoleId, &user.GymId, &user.CreatedAt)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}
	return users, nil
}