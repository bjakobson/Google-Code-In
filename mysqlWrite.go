package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	fmt.Println("GO MySQL Tutorial")

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/testdb")

	if err != nil {
		panic(err.Error())

	}

	defer db.Close()

	insert, err := db.Query("INSERT INTO users (name, lastname) VALUE ('kathy', 'kay')")
	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()
	fmt.Println("Wrote to the database!")

}
