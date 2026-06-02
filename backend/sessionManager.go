package fonction_go

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"log"
	"net/http"
)

var AuthError = errors.New("Unauthorized")

func generateToken(length int) string {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		log.Fatalf("Failed to generate the token :( : %v", err)
	}
	return base64.RawStdEncoding.EncodeToString(bytes)
}

func who(r *http.Request) (string,string){
	var userId string;
	var userName string;
	
	db, err := dbConnection()
	if err != nil {
		println("erreur dans l'ouverture de la DB")
	}

	cookie, err := r.Cookie("session_token")
	if err != nil {
		println("erreur dans l'extraction du cookie")
	}

	err = db.QueryRow(query, cookie.Value).Scan(&userId, &userName)
	if err != nil {
		println("erreur dans l'extraction du pseudo dans la db")
	}
	return userId,userName;
}

func Authorize(r *http.Request) error {
	params_log := extractLog(r)
	db, err := dbConnection()
	isUser, err := userExiste(db, params_log)

	if isUser == false || err != nil {
		return AuthError
	}

	st, err := r.Cookie("session_token")
	if err != nil || st.Value == "" || st.Value != params_log.sessionToken {
		return AuthError
	}

	csrf := r.Header.Get("X-CSRF-Token")
	if csrf != params_log.csrfToken || csrf == "" {
		return AuthError
	}

	return nil
}
