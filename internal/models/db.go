package models

import (
	"database/sql"
		"fmt"
	"os"

		"github.com/joho/godotenv"
_ "github.com/lib/pq"

)

var DB *sql.DB

func InitDB() error {
	_ = godotenv.Load()

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, port)

	var err error

	DB, err = sql.Open("postgres", dsn)

	if err != nil {
		return err
	}

	return DB.Ping()
}