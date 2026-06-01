package fonction_go

import (
	"database/sql"
	"fmt"
)

func insertUser(db *sql.DB, user register) (bool, error) {
	query := `INSERT INTO users (username, email, password) VALUES (?, ?, ?)`
	println("insertion")
	_, err := db.Exec(query, user.username, user.mail, user.password)
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	fmt.Println("User add on the db ;)")
	return true,err
}

func logUser(db *sql.DB, user login) (bool, error) {
	query := `SELECT password FROM users WHERE email=(?)`

	var hashedPassword string;
	err := db.QueryRow(query, user.mail).Scan(&hashedPassword)
	if err != nil {
		return false, err;
	}

	if !checkPassword(hashedPassword, user.password){
		fmt.Println("Bad password")
		return false,nil;
	}

	return true, nil;
}