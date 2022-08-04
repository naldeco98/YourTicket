package main

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/naldeco98/YourTicket/cmd/api/router"
	"github.com/naldeco98/YourTicket/pkg/storage"
)

func main() {
	// Variable definitions
	var (
		err error
		r   = gin.Default()
		dsn = "user1:secret_password@/testYT?charset=utf8&parseTime=True&loc=Local"
	)
	var DB sql.DB

	err = storage.GetMySQL(&DB, dsn)
	if err != nil {
		msg := "error getting database: " + err.Error()
		panic(msg)
	}

	routes := router.NewRouter(r, &DB)
	routes.MapRoutes()

	if err := r.Run("localhost:8080"); err != nil {
		panic(err)
	}
}
