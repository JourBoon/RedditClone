package fonction_go

import (
	"database/sql"
	"fmt"
)

func insertUser(db *sql.DB, user register) (bool, error) {
	query := `INSERT INTO users (username, email, password) VALUES (?, ?, ?)`
	println("insertion")

	passwd := hashedPassword(user.password)
	_, err := db.Exec(query, user.username, user.mail, passwd)
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	fmt.Println("User add on the db ;)")
	return true, nil
}

func postMess(db *sql.DB,subject string, body string) (error){
	query:= `INSERT INTO messages (id_user,subject,body) VALUES (?,?,?)`
	_, err := db.Exec(query, body/*a completer */, subject, body);
	if err != nil {
		fmt.Println(err);
		return err;
	}
	return nil;
}

func logUser(db *sql.DB, user login) (bool, error) {
	query := `SELECT password FROM users WHERE email=(?)`

	var hashedPassword string
	err := db.QueryRow(query, user.mail).Scan(&hashedPassword)
	if err != nil {
		return false, err
	}

	if !checkPassword(hashedPassword, user.password) {
		fmt.Println("Bad password")
		return false, nil
	}

	return true, nil
}