package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ggarber42/go-crud-api/initializers"
	"github.com/ggarber42/go-crud-api/models"
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

func FindAllTodos(context *gin.Context) {
	var todos []models.Todo
	initializers.DB.Find(&todos)
	context.JSON(http.StatusOK, gin.H{"data": todos})
}

func FindTodoById(context *gin.Context) {
	var todo models.Todo

	if err := initializers.DB.Where("id = ?", context.Param("id")).First(&todo).Error; err != nil {
		context.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": todo})
}

type UpdatePostInput struct {
	Done   bool `json:"done"`
}

func UpdateTodo(context *gin.Context) {
	var todo models.Todo
	var input UpdatePostInput
	
	if err := initializers.DB.Where("id = ?", context.Param("id")).First(&todo).Error; err != nil {
		context.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	if err := context.ShouldBindJSON(&input); err != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	initializers.DB.Model(&todo).Update("Done", input.Done)
	context.JSON(http.StatusOK, gin.H{"data": todo})
}

func DeleteTodo(context *gin.Context) {
	var todo models.Todo

	if err := initializers.DB.Where("id = ?", context.Param("id")).First(&todo).Error; err != nil {
		context.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	initializers.DB.Delete(&todo)
	context.JSON(http.StatusNoContent, gin.H{"data": "todo deleted!"})
}
