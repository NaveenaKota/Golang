package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/sijms/go-ora/v2"
)

var localDB = map[string]string{
	"service":  "xe",
	"username": "system",
	"server":   "localhost",
	"port":     "1521",
	"password": "oracle",
}

func main() {
	// Replace with your Oracle database connection details

	// Create the connection string
	connectionString := "oracle://" + localDB["username"] + ":" + localDB["password"] + "@" + localDB["server"] + ":" + localDB["port"] + "/" + localDB["service"]

	// Open a connection to the Oracle database
	db, err := sql.Open("oracle", connectionString)
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	defer db.Close()

	// Ping the database to check the connection
	err = db.Ping()
	if err != nil {
		log.Fatal("Error pinging the database:", err)
	}

	fmt.Println("Connected to the Oracle database!")
}

/*
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-goracle/goracle"
	"github.com/miekg/pkcs11"
)

func main() {
	dsn := godror.ConnectionParams{
		Username:      "system",
		Password:      "oracle",
		ConnectString: "localhost:1521/XE",
	}.String()

	db, err := sql.Open("oci8", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	fmt.Println("Connection success")

}*/
