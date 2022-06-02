package main

import (
	"github.com/gin-gonic/gin"
	"github.com/naldeco98/YourTicket/cmd/api/router"
	"github.com/naldeco98/YourTicket/pkg/db"
)

func main() {
	// Variable definitions
	var (
		err error
		r   = gin.Default()
		dsn = "user1:secret_password@/testYT?charset=utf8&parseTime=True&loc=Local"
	)

	DB, err := db.GetMySQL(dsn)
	if err != nil {
		msg := "error getting database: " + err.Error()
		panic(msg)
	}

	routes := router.NewRouter(r, DB)
	routes.MapRoutes()

	if err := r.Run("localhost:8080"); err != nil {
		panic(err)
	}
}
