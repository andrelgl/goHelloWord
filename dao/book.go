package dao

import (
	"database/sql"
	"log"

	"github.com/lib/pq"
)

// NewDb initialize a instance from db
func NewDb(url string) (*sql.DB, bool) {
	pqURL, err := pq.ParseURL(url)
	if err != nil {
		log.Fatal(err)
	}
	db, err := sql.Open("postgres", pqURL)

	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return db, true
}
