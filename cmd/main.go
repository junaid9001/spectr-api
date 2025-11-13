package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/test", func(ctx *gin.Context) {
		ctx.String(200, "server is running")
	})

	r.Run(":8080")

}
