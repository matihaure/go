package model

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var con *sql.DB

func Connect() *sql.DB {

	db, err := sql.Open("mysql", "mhaure:Mh..35776251.@tcp(mysql:3306)/mysql")

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connected to the DataBase")
		fmt.Println("Serving...")
	}

	con = db

	return db

}
