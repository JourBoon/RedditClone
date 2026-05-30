package fonction_go

import (
	"database/sql"
	"fmt"
)

func createUserTable(db *sql.DB) error {
    query := `
    CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        username TEXT UNIQUE NOT NULL,
        email TEXT UNIQUE NOT NULL,
		password TEXT UNIQUE NOT NULL,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    );`

    fmt.Println("User Table creat ;)")

    _, err := db.Exec(query)
    return err
}