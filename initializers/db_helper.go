package initializers

import (
	"fmt"
	"log"
	"os"

	"github.com/ggarber42/go-crud-api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func connectDB() {
	var err error
	connection := fmt.Sprintf("host=%s user=%s password=%s dbname=%s", os.Getenv("HOST"), os.Getenv("USER"), os.Getenv("PASSWORD"), os.Getenv("DB_NAME"))
	dsn := connection + " " + "port=5432 sslmode=disable timezone=Asia/Shanghai"
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
