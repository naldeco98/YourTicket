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

func NewHandler(u users.Service) *User {
	return &User{user: u}
}

func (u *User) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user domain.User
		if err := c.ShouldBindJSON(&user); err != nil {
			web.Failure(c, 400, err.Error())
			return
		}
		id, err := u.user.Create(c, &user)
		if err != nil || id == 0 {
			web.Failure(c, 409, "username already used")
			return
		}
		web.Success(c, 201, id)
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

func (u *User) Update() gin.HandlerFunc {
	type UserUpdate struct {
		Id       int    `json:"id,omitempty"`
		Username string `json:"username"`
		Password string `json:"password"`
		RoleId   int    `json:"role_id"`
		GymId    int    `json:"gym_id"`
	}
	return func(c *gin.Context) {
		var user UserUpdate
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			web.Failure(c, 400, "wrong id")
		}
		user.Id = id
		if err := c.ShouldBindJSON(&user); err != nil {
			web.Failure(c, 400, err.Error())
			return
		}
		userUpdated, err := u.user.Update(c, &user)
		if err != nil {
			web.Failure(c, 404, err.Error())
			return
		}
		web.Success(c, 200, userUpdated)
	}
}
