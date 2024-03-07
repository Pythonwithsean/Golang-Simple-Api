package main

import (
	"log"

	"github.com/go-sql-driver/mysql"
)

func main() {
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "admin",
		Addr:                 "",
		DBName:               "projectmanager",
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	sqlStorage := NewMySQLStroeage(cfg)
	db, err := sqlStorage.Init()
	if err != nil {
		log.Fatal(err)
	}
	store := NewStore(db)
	api := NewAPISEVER(":8080", store)

}
