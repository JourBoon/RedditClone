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
	defer db.Close()

	params_log.sessionToken = generateToken(32)
	//params_log.CsrfToken = generateToken(32)
	println("InitStart:")
	http.SetCookie(w, &http.Cookie{
		Name:    "username",
		Value:   params_log.username,
		Path:    "/",
		Expires: time.Now().Add(24 * time.Hour),
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    params_log.sessionToken,
		Path:     "/",
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

	if err != nil {
		println("pas de cookie")
		println(err)
		return
	}
	db, err := dbConnection()
	if err != nil {
		handleError(w, "Erreur DB", http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	sessionToken, err := returnSessionToken(db, params_log)

	fmt.Println("Session Token :", sessionToken)
	fmt.Println("Session Token en cookies", st.Value)
	println("Init:", params_log.username)
	if sessionToken != st.Value {
		params_log.sessionToken = generateToken(32)
		//params_log.CsrfToken = generateToken(32)

		http.SetCookie(w, &http.Cookie{
			Name:    "username",
			Value:   params_log.username,
			Path:    "/",
			Expires: time.Now().Add(24 * time.Hour),
		})

		http.SetCookie(w, &http.Cookie{
			Name:     "session_token",
			Value:    params_log.sessionToken,
			Path:     "/",
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
	db, err := dbConnection()
	if err != nil {
		return AuthError
	}
	defer db.Close()

	st, err := r.Cookie("session_token")
	if err != nil {
		return AuthError
	}

	user_cookie, err := r.Cookie("username")
	if err != nil {
		return AuthError
	}

	User, err := returnUsername(db, st.Value)
	if err != nil {
		fmt.Printf("User not in the db")
		return AuthError
	}

	if User == "" || User != user_cookie.Value {
		fmt.Printf("Session Token invalid")
		return AuthError
	}

	return nil
}

func Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "username",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "csrf_token",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		Expires:  time.Unix(0, 0),
		HttpOnly: false,
	})

	params_log := extractLog(r)

	params_log.sessionToken = ""
	params_log.CsrfToken = ""

	fmt.Println("Logged is good ;)")
}
