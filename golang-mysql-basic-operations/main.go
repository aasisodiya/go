package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("MySQL with Go Lang")

	// Open our database connection.
	db, err := sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/testschema")
	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	// Open may just validate its arguments without creating a connection
	// to the database. To verify that the data source name is valid, call
	// Ping.
	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	// defer the close till after the main function has finished
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	tableName := "Persons"
	fname := "Akash"
	lname := "Sisodiya"
	ID := 1

	statement := "INSERT INTO " + tableName + " VALUES( ?, ?, ? )"
	// Prepare statement for inserting data, this will depend on your table
	stmtIns, err := db.Prepare(statement) // ? = placeholder
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtIns.Close() // Close the statement when we leave main() / the program terminates

	_, err = stmtIns.Exec(ID, lname, fname) // Insert the record using statement
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	rows, err := db.Query("SELECT * FROM Persons")
	// above statement will fetch records in rows variable
	defer rows.Close()
	// if there is an error here, handle it
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	var PersonID int
	var LastName string
	var FirstName string
	for rows.Next() {
		// for each row, scan the result into defined variables
		err = rows.Scan(&PersonID, &LastName, &FirstName)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		fmt.Printf("PersonID: %d, FirstName: %s, LastName: %s \n", PersonID, FirstName, LastName)
	}
}
