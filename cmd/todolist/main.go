package main

import (
	"os"

	"github.com/YukiHime23/go-todolist/internal/server"
	"github.com/gin-gonic/gin"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			os.Exit(1)
		}
	}()

	router := gin.Default()

	server.RegisterRoutes(router)

	router.Run(":8080")
}

func Server() {

}
