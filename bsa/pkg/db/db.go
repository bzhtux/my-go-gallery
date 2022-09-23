package db

import (
	"fmt"
	"io/ioutil"
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
	// wokeignore:rule=disable
	dsn    = fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DBHost, DBPort, DBUser, DBName, DBPassword)
	B_DIR  = "/bindings"
	B_NAME = "psql"
	// B_NAME       = os.Getenv("BIND_NAME")
	// B_DIR       = os.Getenv("BIND_DIR")
)

// type Handler struct {
// 	DB *gorm.DB
// }

// func New(db *gorm.DB) Handler {
// 	return Handler{db}
// }

func GetFileContent(f string) string {
	c, err := ioutil.ReadFile(f)
	if err != nil {
		fmt.Println(err)
		return "null"
	}
	fmt.Printf("*** Content of %v : %v", f, c)
	return string(c)
}

func OpenDB() *gorm.DB {

	if DBUser == "" {
		DBUser = GetFileContent(B_DIR + B_NAME + "/username")
	} else {
		fmt.Println("OpenDB.DBUser: " + DBUser)
	}
	if DBName == "" {
		DBName = GetFileContent(B_DIR + B_NAME + "/database")
	} else {
		fmt.Println("OpenDB.DBName: " + DBName)
	}
	if DBHost == "" {
		DBHost = GetFileContent(B_DIR + B_NAME + "/database")
	} else {
		fmt.Println("OpenDB.DBHost: " + DBHost)
	}
	if DBPort == "" {
		DBPort = GetFileContent(B_DIR + B_NAME + "/port")
	} else {
		fmt.Println("OpenDB.DBPort: " + DBPort)
	}
	if DBPassword == "" {
		DBPassword = GetFileContent(B_DIR + B_NAME + "/password")
	} else {
		fmt.Println("OpenDB.DBPassword: " + DBPassword)
	}

	// wokeignore:rule=disable
	dsn = fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DBHost, DBPort, DBUser, DBName, DBPassword)

	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		// log.Fatalf("%s", err)
		fmt.Println("*** Error connectinng to DB ...")
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
