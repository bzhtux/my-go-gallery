package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DBUser       = os.Getenv("DB_USER")
	DBName       = os.Getenv("DB_NAME")
	DBHost       = os.Getenv("DB_HOST")
	DBPort       = os.Getenv("DB_PORT")
	DBPassword   = os.Getenv("DB_PASSWORD")
	SMTPUSer     = os.Getenv("SMTP_USER")
	SMTPPassword = os.Getenv("SMTP_PASSWORD")
	SMTPHost     = os.Getenv("SMTP_HOST")
	SMTPPort     = os.Getenv("SMTP_PORT")
	dsn          = fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DBHost, DBPort, DBUser, DBName, DBPassword)
)

func OpenDB() *gorm.DB {

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("%s", err)
	}

	return db
}

func HealthCheck(dsn string) bool {
	_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return false
	}

	return true
}

func AutoMigrate(db *gorm.DB, database interface{}) {

	db.AutoMigrate(database)

}
