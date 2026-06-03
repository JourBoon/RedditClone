package fonction_go

import (
	"database/sql"
	"fmt"
)

func insertUser(db *sql.DB, user register) (error) {
	query := `INSERT INTO users (username, email, password) VALUES (?, ?, ?)`
	println("insertion")

	passwd := hashedPassword(user.password)
	_, err := db.Exec(query, user.username, user.mail, passwd)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("User add on the db ;)")
	return err
}

func postMess(db *sql.DB,subject string,tag []string ,body string) (error){
	query:= `INSERT INTO messages (id_user,subject,tag,body) VALUES (?,?,?,?)`
	_, err := db.Exec(query, string(tabStringToJson(tag)), subject, body);
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

func getMess(db *sql.DB) ([]mess, error) {
	const postsQuery = `SELECT users.username, messages.subject, messages.body, messages.tag, messages.created_at FROM messages JOIN users ON messages.id_user=users.id`
	rows, err := db.Query(postsQuery)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	rowsData := []mess{}
	for rows.Next() {
		var username string
		var subject string
		var body string
		var tagStr string
		var tag []string
		var created_at string
		err = rows.Scan(&username, &subject, &body, &tagStr, &created_at)

		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		tag = jsonToTabString([]byte(tagStr))

		rowsData = append(rowsData, mess{
			username:   username,
			subject:    subject,
			body:       body,
			tag:        tag,
			created_at: created_at,
		})
	}

	if err := rows.Err(); err != nil {
		fmt.Println(err)
		return rowsData, err
	}

	return rowsData, nil
}
