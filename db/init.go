package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/joho/godotenv"
)

var Connection *gorm.DB

func InitializeDatabase() {
	dbUsername, dbPassword, dbHost, dbPort, dbName := loadEnvVars()
	db := createDbConnection(dbUsername, dbPassword, dbHost, dbPort, dbName)
	migrate(db)
}

func migrate(db *gorm.DB) {
	err := db.AutoMigrate(&MdContent{})
	if err != nil {
		log.Fatal("Unable to migrate database")
	}

}

func createDbConnection(dbUsername, dbPassword, dbHost, dbPort, dbName string) *gorm.DB {
	connectionString := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", dbUsername, dbPassword, dbHost, dbPort, dbName)
	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	Connection = db // Assign global var so other files can access connection
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	return db
}

func loadEnvVars() (dbUsername, dbPassword, dbHost, dbPort, dbName string) {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Read environment variables
	dbUsername = os.Getenv("DB_USERNAME")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbHost = os.Getenv("DB_HOST")
	dbPort = os.Getenv("DB_PORT")
	dbName = os.Getenv("DB_NAME")
	return
}
