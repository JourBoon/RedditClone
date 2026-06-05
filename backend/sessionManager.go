package fonction_go

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"
)

var AuthError = errors.New("Unauthorized")

func generateToken(length int) string {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		log.Fatalf("Failed to generate the token :( : %v", err)
	}
	return base64.RawStdEncoding.EncodeToString(bytes)
}

func InitStartSession(w http.ResponseWriter, r *http.Request) {
	params_log := extractLog(r)

	db, err := dbConnection()
	if err != nil {
		handleError(w, "Erreur DB", http.StatusInternalServerError, err)
		return
	}

	params_log.sessionToken = generateToken(32)
	//params_log.CsrfToken = generateToken(32)

	http.SetCookie(w, &http.Cookie{
		Name:     "username",
		Value:    params_log.username,
		Expires:  time.Now().Add(24 * time.Hour),
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    params_log.sessionToken,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
	})

	/*http.SetCookie(w, &http.Cookie{
		Name:     "csrf_token",
		Value:    params_log.CsrfToken,
		Path:     "/",
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: false,
	})*/

	addSessionToken(db, params_log)
	//addCsrfToken(db, params_log)

}

func InitSession(w http.ResponseWriter, r *http.Request) {
	params_log := extractLog(r)
	st, err := r.Cookie("session_token")

	db, err := dbConnection()
	if err != nil {
		handleError(w, "Erreur DB", http.StatusInternalServerError, err)
		return
	}

	sessionToken, err := returnSessionToken(db, params_log)

	fmt.Println("Session Token :", sessionToken)
	fmt.Println("Session Token en cookies", st.Value)

	if sessionToken != st.Value {
		params_log.sessionToken = generateToken(32)
		//params_log.CsrfToken = generateToken(32)

		http.SetCookie(w, &http.Cookie{
			Name:     "username",
			Value:    params_log.username,
			Expires:  time.Now().Add(24 * time.Hour),
		})

		http.SetCookie(w, &http.Cookie{
			Name:     "session_token",
			Value:    params_log.sessionToken,
			Expires:  time.Now().Add(24 * time.Hour),
			HttpOnly: true,
		})

		/*http.SetCookie(w, &http.Cookie{
			Name:     "csrf_token",
			Value:    params_log.CsrfToken,
			Path:     "/",
			Expires:  time.Now().Add(24 * time.Hour),
			HttpOnly: false,
		})*/

		addSessionToken(db, params_log)
		//addCsrfToken(db, params_log)
	} else {
		return
	}
}

func Authorize(r *http.Request) error {
	params_log := extractLog(r)
	db, err := dbConnection()
	st, err := r.Cookie("session_token")
	user_cookie, err := r.Cookie("username")

	User, err := returnUsername(db, st.Value)

	if User != user_cookie.Value || err != nil {
		fmt.Printf("User not in the db")
		return AuthError
	}

	tokenSession, err := returnSessionToken(db, params_log)

	if err != nil || st.Value == "" || st.Value != tokenSession {
		fmt.Printf("Session Token invalid")
		return AuthError
	}

	// Protection du jeton de session contre les attaques csrf a revoir

	/*
	csrfToken, err := returnCsrfToken(db, params_log)
	csrf := r.Header.Get("X-CSRF-Token")

	fmt.Println("csrf reçu :", csrf)
	fmt.Println("csrf attendu :", csrfToken)

	if csrf != csrfToken || csrf == "" {
		fmt.Println("Csrf Token invalid")
		return AuthError
	}*/

	return nil
}

func Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "username",
		Value:    "",
		Path:     "/",
		MaxAge: -1,
		Expires: time.Unix(0, 0),
		HttpOnly: true,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    "",
		Path:     "/",
		MaxAge: -1,
		Expires: time.Unix(0, 0),
		HttpOnly: true,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "csrf_token",
		Value:    "",
		Path:     "/",
		MaxAge: -1,
		Expires: time.Unix(0, 0),
		HttpOnly: false,
	})

	params_log := extractLog(r)

	params_log.sessionToken = ""
	params_log.CsrfToken = ""


	fmt.Println("Logged is good ;)")
}
