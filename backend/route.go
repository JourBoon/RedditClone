package fonction_go

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

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

func LogoutBtn(w http.ResponseWriter, r *http.Request) {
	path := "index.html"
	renderTemplateWithData(w, path, nil)
	Logout(w, r)
}

func Home(w http.ResponseWriter, r *http.Request) {

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
			InitStartSession(w, r)
			path := "/protected/home.html"
			data,err := getMess(db)
			if err != nil {
				handleError(w, "Erreur dans le chargement des messages", http.StatusInternalServerError, err)
				return
			}	
			renderTemplateWithData(w, path,data )
			defer db.Close()
			return
		}
	}

	path := "/auth/login.html"
	renderTemplateWithData(w, path, nil)
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
	defer db.Close()

	username, err := getUsernameFromSessionCookie(r)
	if err != nil {
		handleError(w, "Utilisateur non authentifié", http.StatusUnauthorized, err)
		return
	}

	params_mess := extractMess(r)
	user_id := getIdUserByUsername(db, username)
	postMess(db, user_id, params_mess.subject, params_mess.tags, params_mess.body)
	path := "/protected/create.html"
	renderTemplateWithData(w, path, nil)
}

func RoutePages() {
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

func getUsernameFromSessionCookie(r *http.Request) (string, error) {
	st, err := r.Cookie("session_token")
	if err != nil {
		return "", err
	}

	db, err := dbConnection()
	if err != nil {
		return "", err
	}
	defer db.Close()

	return returnUsername(db, st.Value)
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

	username, err := getUsernameFromSessionCookie(r)
	if err != nil {
		fmt.Println("Unable to resolve username from session cookie")
		return
	}
	fmt.Printf("CSRF validate ;) Welcome to the forum %s", username)
}
