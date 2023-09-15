package main

import (
	"fmt"

	"github.com/ggarber42/go-crud-api/initializers"
	"github.com/ggarber42/go-crud-api/models"
)

func init() {
	initializers.ConnectDB()
}

func main() {
	initializers.DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	initializers.DB.AutoMigrate(&models.Todo{})
	fmt.Println("👍 Migration complete")
}