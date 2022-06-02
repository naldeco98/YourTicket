package router

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/naldeco98/YourTicket/cmd/api/handler"
	"github.com/naldeco98/YourTicket/internal/users"
)

type Router interface {
	MapRoutes()
}

type router struct {
	r  *gin.Engine
	rg *gin.RouterGroup
	db *sql.DB
}

func NewRouter(r *gin.Engine, db *sql.DB) Router {
	return &router{r: r, db: db}
}

func (r *router) MapRoutes() {
	r.setGroup()
	r.healthCheck()
	r.buildUsersRoutes()
}

func (r *router) setGroup() {
	r.rg = r.r.Group("/api/v1")
}

func (r *router) healthCheck() {
	r.r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
}
func (r *router) buildUsersRoutes() {

	u := r.rg.Group("/users")
	repo := users.NewRepository(r.db)
	service := users.NewService(repo)
	handler := handler.NewHandler(service)
	{
		u.POST("", handler.Register())
	}
}
