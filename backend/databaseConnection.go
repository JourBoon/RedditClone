package fonction_go

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func dbConnection() (*sql.DB, error) {
    db, err := sql.Open("sqlite3", "./database/database.sqlite")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    err = db.Ping()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Successfully connected to SQLite database!")

    return db, nil
}