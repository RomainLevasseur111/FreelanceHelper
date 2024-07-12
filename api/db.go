package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func createDB(){
	db, err := sql.Open(DRIVER, DB)
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()
	request := `
	CREATE TABLE IF NOT EXISTS USERS (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		name VARCHAR(20),
		password TEXT
	);
	CREATE TABLE IF NOT EXISTS SIMULATION (
		user_id INTEGER NOT NULL,
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		status TEXT NOT NULL,
		daily_rate INTEGER,
		monthly_rate INTEGER,
		contract_length VARCHAR(20) NOT NULL,
		daily_charges INTEGER DEFAULT 0,
		monthly_charges INTEGER DEFAULT 0
	);

	`

	_,err = db.Exec(request)
	if err != nil{
		log.Println("DATABASE CREATION ERROR")
		fmt.Println(err)
	}
}