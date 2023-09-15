package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/ggarber42/go-crud-api/controllers"
	"github.com/ggarber42/go-crud-api/initializers"
)

func init(){
	initializers.ConnectDB()
}

func main(){
	server := gin.Default()

	server.GET("/hello", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message": "go crud api",
		});
	})

	server.POST("/todos", controllers.CreateTodo)



	server.Run("localhost:5000");
	fmt.Printf("go crud log")
}