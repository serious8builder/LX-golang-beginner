package main

/*
https://github.com/mattn/go-sqlite3/blob/master/_example/simple/simple.go
*/

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

import (
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	os.Remove("./foo.db")
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := `
	CREATE TABLE FOO (
	id INTEGER NOT NULL PRIMARY KEY,
	name TEXT);
	`

	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		return
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := tx.Prepare("INSERT INTO foo(id, name) values(?, ?)")
	if err != nil {
		log.Fatal(err)
	}

	defer stmt.Close()

	for i := 0; i < 10; i++ {
		_, err = stmt.Exec(i+1, fmt.Sprintf("Hello world [%03d]", i+1))
		if err != nil {
			log.Fatal(err)
		}
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT id, name from foo")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ROW(id=%v, name=%v)\n", id, name)
	}

	os.Remove("./foo.db")
	fmt.Println("END")
}
