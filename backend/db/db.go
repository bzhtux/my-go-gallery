package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenDB(dsn string) *gorm.DB {

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
