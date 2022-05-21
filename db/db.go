package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
n
func OpenDB(dsn string) *gorm.DB {

	// databaseDriver := os.Getenv("DATABASE_DRIVER")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("%s", err)
	}
	// if err := AutoMigrate(db); err != nil {
	// 	panic(err)
	// }
	return db
}

func AutoMigrate(db *gorm.DB, database interface{}) error {

	return db.AutoMigrate(db, database).Error

}
