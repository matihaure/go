package model

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var con *sql.DB

func Connect() *sql.DB{

	db,err := sql.Open("mysql","mhaure:Mh..35776251.@tcp(localhost:3306)/mysql")

	if err!=nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to the DataBase")
	fmt.Println("Serving...")

	con = db

	return db

}