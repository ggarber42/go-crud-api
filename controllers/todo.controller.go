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
	initializers.DB.Create(&todo)

	context.JSON(http.StatusOK, gin.H{"data": todo})
}