package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/naldeco98/YourTicket/internal/domain"
	"github.com/naldeco98/YourTicket/internal/users"
	"github.com/naldeco98/YourTicket/pkg/web"
)

type User struct {
	user users.Service
}

type requestUser struct {
	Id       int    `json:"-"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	RoleId   int    `json:"role_id,omitempty"`
	GymId    int    `json:"gym_id,omitempty"`
}

func NewHandler(u users.Service) *User {
	return &User{user: u}
}

func (u *User) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req requestUser
		if err := c.ShouldBindJSON(&req); err != nil {
			web.Failure(c, 400, err.Error())
			return
		}
		id, err := u.user.Create(c, req.Username, req.Password, req.RoleId, req.GymId)
		if err != nil || id == 0 {
			web.Failure(c, 409, err.Error())
			return
		}
		req.Id = id
		web.Success(c, 201, req)
	}
}

func (u *User) GetById() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Failure(c, 400, "wrong id")
		}
		user, err := u.user.GetById(c, id)
		if err != nil {
			web.Failure(c, 404, err.Error())
			return
		}
		web.Success(c, 200, user)
	}
}

func (u *User) GetAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		users, err := u.user.GetAll(c)
		if err != nil {
			web.Failure(c, 409, err.Error())
			return
		}
		web.Success(c, 200, users)
	}
}

func (u *User) Put() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			user domain.User
			err  error
		)
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Failure(c, 400, "wrong id")
		}
		user.Id = id
		if err = c.ShouldBindJSON(&user); err != nil {
			web.Failure(c, 400, err.Error())
			return
		}
		if err = web.VerifyEmptyFields(user); err != nil {
			web.Failure(c, 422, err.Error())
			return
		}
		userUpdated, err := u.user.Update(c, user)
		if err != nil {
			web.Failure(c, 409, err.Error())
			return
		}
		web.Success(c, 200, userUpdated)
	}
}
