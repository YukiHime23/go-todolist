package server

import (
	"net/http"

	"github.com/YukiHime23/go-todolist/internal/api"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	var APIv1 = router.RouterGroup

	api.RegisterTodo(&APIv1)
}
