package fonction_go

import (
	"database/sql"
	"errors"
	"fmt"
)

func insertUser(db *sql.DB, user register) error {
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

func postMess(db *sql.DB, id_user string, subject string, tag []string, body string) error {
	query := `INSERT INTO messages (id_user,subject,tags,body) VALUES (?,?,?,?)`
	println(query, id_user, subject, string(tabStringToJson(tag)), body)
	_, err := db.Exec(query, id_user, subject, string(tabStringToJson(tag)), body)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func logUser(db *sql.DB, user login) (bool, error) {
	query := `SELECT password FROM users WHERE email=(?)`

	var hashedPassword string
	err := db.QueryRow(query, user.mail).Scan(&hashedPassword)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil
		}
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

	_, err := db.Exec(query, login.sessionToken, login.mail)
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	fmt.Println("Token add on the db ;)")
	return true, nil
}

func returnSessionToken(db *sql.DB, user login) (string, error) {
	query := `SELECT session.sessionToken FROM session INNER JOIN users ON session.id_user = users.id WHERE users.email=(?)`

	var sessionToken string
	err := db.QueryRow(query, user.mail).Scan(&sessionToken)
	if err != nil {
		fmt.Println("Error return session token")
	}

	return sessionToken, nil
}

func returnCsrfToken(db *sql.DB, user login) (string, error) {
	query := `SELECT session.csrfToken FROM session INNER JOIN users ON session.id_user = users.id WHERE users.email=(?)`

	var csrfToken string
	err := db.QueryRow(query, user.mail).Scan(&csrfToken)
	if err != nil {
		fmt.Println("Error return csrf token")
	}

	return csrfToken, nil
}

func addCsrfToken(db *sql.DB, login login) (bool, error) {
	query := `INSERT INTO session (id_user, csrfToken) SELECT id, ? FROM users WHERE email = ? ON CONFLICT(id_user) DO UPDATE SET csrfToken = excluded.csrfToken;`

	_, err := db.Exec(query, login.CsrfToken, login.mail)
	if err != nil {
		fmt.Println(err)
		return false, err
	}

	fmt.Println("csrf Token add on the db ;)")
	return true, nil
}

func returnUsername(db *sql.DB, sessionToken string) (string, error) {
	query := `SELECT users.username FROM users INNER JOIN session ON session.id_user = users.id WHERE session.sessionToken=(?)`
	var username string = ""
	err := db.QueryRow(query, sessionToken).Scan(&username)
	if err != nil {
		fmt.Println(err)
	}
	return username, err
}

func getMess(db *sql.DB) ([]Mess, error) {
	const postsQuery = `SELECT messages.id, users.username, messages.id_user, messages.subject, messages.body, messages.tags, messages.likes, messages.created_at FROM messages JOIN users ON messages.id_user=users.id`
	rows, err := db.Query(postsQuery)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	rowsData := []Mess{}
	for rows.Next() {
		var id int
		var username string
		var userId int
		var subject string
		var body string
		var tagStr string
		var tag []string
		var likes int
		var created_at string
		err = rows.Scan(&id, &username, &userId, &subject, &body, &tagStr, &likes, &created_at)

		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		tag = jsonToTabString([]byte(tagStr))

		rowsData = append(rowsData, Mess{
			Id:         id,
			Username:   username,
			UserId:     userId,
			Subject:    subject,
			Body:       body,
			Tag:        tag,
			Likes:      likes,
			Created_at: created_at,
		})
	}

	if err := rows.Err(); err != nil {
		fmt.Println(err)
		return rowsData, err
	}

	return rowsData, nil
}

func getIdUserByUsername(db *sql.DB, userName string) string {
	query := `SELECT id FROM users WHERE username=(?)`
	var id string
	err := db.QueryRow(query, userName).Scan(&id)
	if err != nil {
		fmt.Println(err)
	}
	return id
}

func addLike(db *sql.DB, userId string, messId string) {
	query := `UPDATE message SET likes += 1 WHERE username=()`
	var id string
	err := db.QueryRow(query, userId, messId).Scan(&id)
	if err != nil {
		fmt.Println(err)
	}
}
