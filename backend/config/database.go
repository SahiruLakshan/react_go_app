package config

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/godror/godror"
)

func ConnectDatabase() *sql.DB {

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	connectString := os.Getenv("DB_CONNECT_STRING")

	dsn := user + "/" + password + "@" + connectString

	db, err := sql.Open("godror", dsn)

	if err != nil {
		log.Fatal("Failed to open database:", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Failed to connect to Oracle:", err)
	}

	log.Println("Oracle database connected")

	return db
}