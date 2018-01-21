package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var (
	DbDriver    string
	DbHost      string
	DbPort      string
	DbUsername  string
	DbPassword  string
	DbDatabase  string
	DbCharset   string
	DbCollation string
	DbPrefix    string
	DbEngine    string
)

func mysql() error {

	DbDriver = os.Getenv("DB_DRIVER")
	DbHost = os.Getenv("DB_HOST")
	DbPort = os.Getenv("DB_PORT")
	DbUsername = os.Getenv("DB_USERNAME")
	DbPassword = os.Getenv("DB_PASSWORD")
	DbDatabase = os.Getenv("DB_DATABASE")
	DbCharset = os.Getenv("DB_CHARSET")
	DbCollation = os.Getenv("DB_COLLATION")
	DbPrefix = os.Getenv("DB_PREFIX")
	DbEngine = os.Getenv("DB_ENGINE")

	// test conn
	// root:root@tcp(localhost:3306)/dbname?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=true&loc=PRC
	db, err := sql.Open(DbDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", DbUsername, DbPassword, DbHost, DbPort, DbDatabase))
	if err != nil {
		// panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
		return err
	}
	defer db.Close()

	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		// panic(err.Error()) // proper error handling instead of panic in your app
		return err
	}

	return nil
}
