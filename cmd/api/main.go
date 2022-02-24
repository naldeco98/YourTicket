package main

import (
	"github.com/gin-gonic/gin"
	"github.com/naldeco98/YourTicket/cmd/api/router"
)

func main() {
	r := gin.Default()
	routes := router.NewRouter(r)
	routes.MapRoutes()

	if err := r.Run("localhost:8080"); err != nil {
		panic(err)
	}
}
