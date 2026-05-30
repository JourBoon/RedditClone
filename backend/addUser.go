package fonction_go

import (
	"database/sql"
	"fmt"
)

func insertUser(db *sql.DB, user queryParams) (int64, error) {
	query := `INSERT INTO users (username, email, password) VALUES (?, ?, ?)`

	result, err := db.Exec(query, user.username, user.mail, user.password)
	if err != nil {
		return 0, err
	}

	fmt.Println("User add on the db ;)")
	return result.LastInsertId()
}
