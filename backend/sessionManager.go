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

func Authorize(r *http.Request) error {
	params_log := extractLog(r)

	user, ok := users[params_log.username]
	if !ok {
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