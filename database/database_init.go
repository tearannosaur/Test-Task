package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type DBModule struct {
	Db *sqlx.DB
}

func DBInit() (*DBModule, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("gotodenv load error")
		return nil, err
	}
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")
	connection := fmt.Sprintf("host=db port=5432 user=%s password=%s dbname=%s sslmode=disable", user, password, name)
	database, err := sqlx.Connect("postgres", connection)
	if err != nil {
		log.Println("database connection error")
		return nil, err
	}
	return &DBModule{Db: database}, nil
}

func (d *DBModule) DBClose() error {
	err := d.Db.Close()
	if err != nil {
		log.Println("database close error")
		return err
	}
	return nil
}
