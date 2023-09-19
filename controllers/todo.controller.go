package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"github.com/ggarber42/go-crud-api/models"
	"github.com/ggarber42/go-crud-api/initializers"
)

type CreateTodoInput struct {
	Title   string `json:"title" binding:"required"`
}

func CreateTodo(context *gin.Context) {
	var input CreateTodoInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo := models.Todo{ Title: input.Title }
	context.JSON(http.StatusOK, gin.H{"data": todo})
}

func FindAllTodos(context *gin.Context) {
	var todos []models.Todo
	initializers.DB.Find(&todos)
	context.JSON(http.StatusOK, gin.H{"data": todos})
}

type UpdateTodoInput struct {
	Done bool `json:"done" binding:"required"`
}

	context.JSON(http.StatusOK, gin.H{"data": todo})
}
