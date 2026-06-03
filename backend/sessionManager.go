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

func Logout(w http.ResponseWriter, r *http.Request) {
	if err := Authorize(r); err != nil {
		er := http.StatusUnauthorized
		http.Error(w, "Unauthorized", er)
		return 
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "username",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: true,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: true,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "csrf_token",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HttpOnly: false,
	})

	params_log := extractLog(r)

	params_log.sessionToken = ""
	params_log.csrfToken = ""


	fmt.Println("Logged is good ;)")
}
