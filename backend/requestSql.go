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

func addSessionToken(db *sql.DB, login login) (bool, error) {
	query := `INSERT INTO session (id_user, sessionToken) SELECT id, ? FROM users WHERE email = ? ON CONFLICT(id_user) DO UPDATE SET sessionToken = excluded.sessionToken;`
	// Requête SQL suggérée par ChatGpt ;)

	_, err := db.Exec(query, login.sessionToken, login.mail)
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	fmt.Println("Token add on the db ;)")
	return true, nil
}

func returnSessionToken(db *sql.DB, user login) (bool, error) {
	query := `SELECT session.sessionToken FROM session INNER JOIN users ON session.id_user = users.id WHERE users.email=(?)`

	var sessionToken string
	err := db.QueryRow(query, user.mail).Scan(&sessionToken)
	if err != nil {
		return false, err
	}

	if sessionToken != user.sessionToken {
		fmt.Println("Bad token session")
		return false, nil
	}

	return true, nil
}

func addCsrfToken(db *sql.DB, login login) (bool, error) {
	query := `INSERT INTO session (id_user, csrfToken) SELECT id, ? FROM users WHERE email = ? ON CONFLICT(id_user) DO UPDATE SET csrfToken = excluded.csrfToken;`
	// Requête SQL suggérée par ChatGpt ;)

	_, err := db.Exec(query, login.csrfToken, login.mail)
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	fmt.Println("csrf Token add on the db ;)")
	return true, nil
}

func userExiste(db *sql.DB, login login) (bool, error) {
	query := `SELECT username FROM users WHERE username=(?)`
	var username string = "";
	err := db.QueryRow(query, login.mail).Scan(&username)
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	if (username!=""){
		fmt.Println("User return with sucess ;)")
		return true, nil
	}
	return false,err;
}
