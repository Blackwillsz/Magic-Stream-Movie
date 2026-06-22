package main

import (
	"fmt"

	controller "github.com/Blackwillsz/Magic-Stream-Movie/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/hello", func(ctx *gin.Context) {
		ctx.String(200, "Hello, Magic Steam!")
	})

	router.GET("/movies", controller.GetMovies())

	if err := router.Run(":8080"); err != nil {
		fmt.Println("Fail to start server", err)
	}
}
