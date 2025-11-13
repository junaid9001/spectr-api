package main

import (
	"github.com/gin-gonic/gin"
	"github.com/junaid9001/spectr-api/config"
	"github.com/junaid9001/spectr-api/routes"
)

func main() {
	r := gin.Default()

	config.ConnectDB()

	r.GET("/test", func(ctx *gin.Context) {
		ctx.String(200, "server is running")
	})

	routes.RegisterRoutes(r)

	r.Run(":8080")

}
