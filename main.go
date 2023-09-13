package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main(){
	server := gin.Default()
	server.GET("/hello", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message": "go crud api",
		});
	})
	server.Run("localhost:5000");
	fmt.Printf("go crud log")
}