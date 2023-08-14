package database

import (
	"database/sql"
	"log"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func Connect_to_DB() *sql.DB{
	db, err := sql.Open("mysql", "root:Naveena@123@tcp(127.0.0.1:3306)/library_DB")
	if err != nil {
		log.Fatal(err)
	}
	// Test the connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected to database!")
	return db
}
