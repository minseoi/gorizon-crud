package db

import (
	"fmt"
	"log"
	"os"

	"github.com/minseoi/gorizon/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Initialize() {
	var err error

	// MySQL connection string from environment variables or use defaults
	dbUser := getEnv("DB_USER", "root")
	dbPassword := getEnv("DB_PASSWORD", "password")
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "3306")
	dbName := getEnv("DB_NAME", "gorizon_db")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	log.Println("Database connection established successfully")

	// Auto migrate the schema
	err = DB.AutoMigrate(&models.Ingredient{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	log.Println("Database migration completed successfully")

	// Seed initial data if table is empty
	var count int64
	DB.Model(&models.Ingredient{}).Count(&count)
	if count == 0 {
		seedData()
	}
}

func seedData() {
	ingredients := []models.Ingredient{
		{Name: "토마토"},
		{Name: "양파"},
		{Name: "당근"},
	}

	for _, ingredient := range ingredients {
		if err := DB.Create(&ingredient).Error; err != nil {
			log.Printf("Failed to seed ingredient %s: %v", ingredient.Name, err)
		}
	}

	log.Println("Initial data seeded successfully")
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
