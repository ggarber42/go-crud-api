package main

import (
	"os"

	"github.com/gin-gonic/gin"

	"github.com/ggarber42/go-crud-api/controllers"
	"github.com/ggarber42/go-crud-api/initializers"
)

func init(){
	initializers.Run()
}

func main(){
	server := gin.Default()
	port := os.Getenv("PORT")

	server.GET("/hello", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message": "go crud api",
		});
	})

	server.POST("/todos", controllers.CreateTodo)
	server.GET("/todos", controllers.FindAllTodos)
	server.GET("/todos/:id", controllers.FindTodoById)
	server.PUT("/todos/:id", controllers.UpdateTodo)
	server.DELETE("/todos/:id", controllers.DeleteTodo)

	server.Run(port)
}
