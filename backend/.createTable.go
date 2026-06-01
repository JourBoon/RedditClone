package fonction_go

import (
	"database/sql"
	"fmt"
)

func createUserTable(db *sql.DB) (bool, error) {
	query := `
    CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        username TEXT UNIQUE NOT NULL,
        email TEXT UNIQUE NOT NULL,
		password TEXT UNIQUE NOT NULL,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP
    );`

    _, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	fmt.Println("Users table creat :)")
	return true,err
}