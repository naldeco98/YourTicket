package handler

import (
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

func (u *User) Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user domain.User
		if err := c.ShouldBindJSON(&user); err != nil {
			web.Failure(c, 400, err.Error())
			return
		}
		id, err := u.user.Register(c, &user)
		if err != nil {
			web.Failure(c, 409, err.Error())
			return
		}
		web.Success(c, 201, id)
	}
}
