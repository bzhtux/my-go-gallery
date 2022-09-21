package db

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

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
	B_DIR        = "/bindings"
)

// type Handler struct {
// 	DB *gorm.DB
// }

// func New(db *gorm.DB) Handler {
// 	return Handler{db}
// }

func GetFileContent(f string) string {
	c, err := ioutil.ReadFile(f)
	// c, err := io.READ
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}
	return string(c)
}

func OpenDB() *gorm.DB {

	for {
		time.Sleep(1000)
		_, err := os.Stat(B_DIR + "/psql")
		if err != nil {
			break
		} else {
			// fmt.Println(B_DIR + "/psql does not exists ... waiting ...")
			log.Printf(err)
		}
	}

	if DBUser == "" {
		os.Setenv(DBUser, GetFileContent(B_DIR+"/psql/username"))
	} else {
		fmt.Println("OpenDB.DBUser: " + DBUser)
	}
	if DBName == "" {
		os.Setenv(DBName, GetFileContent(B_DIR+"/psql/database"))
	} else {
		fmt.Println("OpenDB.DBName: " + DBName)
	}
	if DBHost == "" {
		os.Setenv(DBHost, GetFileContent(B_DIR+"/psql/host"))
	} else {
		fmt.Println("OpenDB.DBHost: " + DBHost)
	}
	if DBPort == "" {
		os.Setenv(DBPort, GetFileContent(B_DIR+"/psql/port"))
	} else {
		fmt.Println("OpenDB.DBPort: " + DBPort)
	}
	if DBPassword == "" {
		os.Setenv(DBPassword, GetFileContent(B_DIR+"/psql/password"))
	} else {
		fmt.Println("OpenDB.DBPassword: " + DBPassword)
	}

	dsn = fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DBHost, DBPort, DBUser, DBName, DBPassword)

	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		// log.Fatalf("%s", err)
		log.Printf("%s", err)
	}

	return conn
}

func HealthCheck(dsn string) bool {
	_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return false
	} else {
		return true
	}
}

func AutoMigrate(db *gorm.DB, database interface{}) {

	db.AutoMigrate(database)

}
