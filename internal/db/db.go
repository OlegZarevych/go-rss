package db

import (
	"database/sql"
	"fmt"
	xml "go-rss/internal/xml"

	_ "github.com/lib/pq"
)

const (
	host     = "postgres"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "rss"
	driver   = "postgres"
)

func ConnectDb() (*sql.DB, error) {

	db, err := sql.Open(driver, fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname))
	if err != nil {
		return nil, err
	}
	db.Ping()

	return db, nil
}

func CreateTable(db *sql.DB) {
	fmt.Println("Create table")
	db.QueryRow("CREATE TABLE records (id SERIAL PRIMARY KEY,xml TEXT);")
}

func CreateRecord(db *sql.DB, item xml.Item) (int, error) {
	var id int
	err := db.QueryRow("INSERT INTO records(xml) VALUES($1) RETURNING id", fmt.Sprintf("%s %s %s %s %s", item.Title, item.Description, item.Link, item.PubDate, item.Guid)).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
