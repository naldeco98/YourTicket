package users

import (
	"database/sql"

	"github.com/naldeco98/YourTicket/internal/domain"
)

const (
	create          = `INSERT INTO users (username, password, role_id, gym_id, created_at) VALUES (?, ?, ?, ?, ?);`
	read            = `SELECT id, username, role_id, gym_id, created_at FROM users WHERE id = ?;`
	lookForUsername = `SELECT id FROM users WHERE username = ?;`
)

type Repository interface {
	Create(u *domain.User) (int, error)
	LookByUsername(username string) (int, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(user *domain.User) (int, error) {
	res, err := r.db.Exec(create, user.Username, user.Password, user.RoleId, user.GymId, user.CreatedAt)
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
	err := r.db.QueryRow(lookForUsername, username).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
