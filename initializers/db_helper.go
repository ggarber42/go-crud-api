package initializers

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/ggarber42/go-crud-api/models"
)

var DB *gorm.DB

func connectDB() {
	var err error
	dsn := "host=db user=admin password=admin dbname=crud_db port=5432 sslmode=disable timezone=Asia/Shanghai"
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the Database")
	}
	fmt.Println("🚀 Connected Successfully to the Database")
}

func migrate(){
	DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	DB.AutoMigrate(&models.Todo{})
	fmt.Println("👍 Migration complete")
}

func Run() {
	connectDB()
	migrate()
}
