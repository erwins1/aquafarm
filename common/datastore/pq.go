package datastore

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func InitPostgres() *sqlx.DB {
	var (
		db  *sqlx.DB
		err error
	)

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	// user=campaign_user password=BQqVxrJ2aNLnZXm dbname=tokopedia-campaign host=10.21.33.176 port=5432 sslmode=disable
	strConn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		dbUser,
		dbPassword,
		dbName,
		dbHost,
		dbPort)

	if strConn == "" {
		return nil
	}

	db, err = sqlx.Connect("postgres", strConn)
	if err != nil {
		log.Println(err, " =====")
		return nil
	}
	return db
}
