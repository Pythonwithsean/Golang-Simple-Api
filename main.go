package main

import (
	"log"
)

func main() {


	sqlStorage := NewMySQLStroeage()
	db, err := sqlStorage.Init()
	if err != nil {
		log.Fatal(err)
	}
	store := NewStore(db)
	api := NewAPISEVER(":8080",store)
	api.Serve()


}
