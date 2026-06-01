package fonction_go

import (
	"database/sql"
	"fmt"
)

func insertUser(db *sql.DB, user register) (int64, error) {
	query := `INSERT INTO users (username, email, password) VALUES (?, ?, ?)`

	result, err := db.Exec(query, user.username, user.mail, user.password)
	if err != nil {
		return 0, err
	}

	fmt.Println("User add on the db ;)")
	return result.LastInsertId()
}

func logUser(db *sql.DB, user login) (bool, error) {
	query := `SELECT password FROM users WHERE email=(?)`

	var hashedPassword string;
	err := db.QueryRow(query, user.mail_log).Scan(&hashedPassword)
	if err != nil {
		return false, err;
	}

	if !checkPassword(hashedPassword, user.password_log){
		fmt.Println("Bad password")
		return false,nil;
	}

	return true, nil;
}