package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type todoImpl struct{}

func RegisterTodo(router *gin.RouterGroup) {
	impl := &todoImpl{}

	router.GET("/todos", impl.todoList)
}

func (impl *todoImpl) todoList(c *gin.Context) {
	c.JSON(http.StatusOK, "test")
}
