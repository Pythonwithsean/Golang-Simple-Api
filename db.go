package main

import (
	"database/sql"
	"log"
)

type MySQLStorage struct {
	db *sql.DB
}

func NewMySQLStroeage() *MySQLStorage {
	db, err := sql.Open("mysql", "root:password@tcp(mysql-container:3306)/mydatabase")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MYSQL")
	return &MySQLStorage{db: db}

}

func (s *MySQLStorage) Init() (*sql.DB, error) {

	return s.db, nil
}
