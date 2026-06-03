package fonction_go

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func InitSession(w http.ResponseWriter, r *http.Request) {
	params_log := extractLog(r)

	db, err := dbConnection()
	if err != nil {
		handleError(w, "Erreur DB", http.StatusInternalServerError, err)
		return
	}

	sessionToken, err := returnSessionToken(db, params_log)
	if sessionToken == true {
		params_log.sessionToken = generateToken(32)
		params_log.csrfToken = generateToken(32)

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

		http.SetCookie(w, &http.Cookie{
			Name:     "csrf_token",
			Value:    params_log.csrfToken,
			Expires:  time.Now().Add(24 * time.Hour),
			HttpOnly: false,
		})

		addSessionToken(db, params_log)
		addCsrfToken(db, params_log)
		}
}

func Start(w http.ResponseWriter, r *http.Request) {
	renderTemplateWithData(w, "index.html", nil)
}

func Login(w http.ResponseWriter, r *http.Request) {
	path := "/auth/login.html"
	renderTemplateWithData(w, path, nil)
}

func Register(w http.ResponseWriter, r *http.Request) {
	db, err := dbConnection()
	if err != nil {
		handleError(w, "Erreur DB", http.StatusInternalServerError, err)
		return
	}
	params_reg := extractReg(r)
	erro := insertUser(db, params_reg)
	if erro != nil {
		handleError(w, "Erreur lors de l'insertion dans la database", http.StatusInternalServerError, err)
		return
	}
	path := "/auth/register.html"
	renderTemplateWithData(w, path, nil)
}

func LogoutBtn(w http.ResponseWriter,r *http.Request) {
	Logout(w, r)
	path := "index.html"
	renderTemplateWithData(w, path, nil)
}

func Home(w http.ResponseWriter, r *http.Request) {
	InitSession(w, r)

	db, err := dbConnection()
	if err != nil {
		handleError(w, "Erreur DB", http.StatusInternalServerError, err)
		return
	}

	if r.Method == http.MethodPost {
		params_log := extractLog(r)
		l, err := logUser(db, params_log)
		if err != nil {
			handleError(w, "Erreur lors de la connexion", http.StatusInternalServerError, err)
			return
		}
		if l {
			path := "/protected/home.html"
			renderTemplateWithData(w, path, nil)
			if err != nil {
				handleError(w, "Erreur lors de l'insertion des token dans la db", http.StatusInternalServerError, err)
				return
			}
		}
	}
	defer db.Close()
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	path := "/protected/home.html"
	renderTemplateWithData(w, path, nil)
}

func createForum(w http.ResponseWriter, r *http.Request) {
	db, err := dbConnection()
	if err != nil {
		handleError(w, "Erreur DB", http.StatusInternalServerError, err)
		return
	}
	params_mess := extractMess(r);
	postMess(db,params_mess.subject,params_mess.tags,params_mess.body);
	path := "/protected/create.html"
	renderTemplateWithData(w, path, nil)
}

func DefaultRoutePages() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	publicPath := filepath.Join(dir, "static")	
	fs := http.FileServer(http.Dir(publicPath))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/", Start)
	http.HandleFunc("/login", Login)
	http.HandleFunc("/register", Register)
	http.HandleFunc("/home", Home)
	http.HandleFunc("/createForum", createForum)
	http.HandleFunc("/LogoutBtn", LogoutBtn)
	http.HandleFunc("/HomePage", HomePage)
}

func Protected(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		er := http.StatusMethodNotAllowed
		http.Error(w, "Invalid request method", er)
		return
	}

	if err := Authorize(r); err != nil {
		er := http.StatusUnauthorized
		http.Error(w, "Unauthorized", er)
		return
	}

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	publicPath := filepath.Join(dir, "static")	
	fs := http.FileServer(http.Dir(publicPath))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.HandleFunc("/home", Home)

	parms_log := extractLog(r);
	username := parms_log.username
	fmt.Printf("CSRF validate ;) Welcome to the forum %s", username)
}