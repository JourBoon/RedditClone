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
	path := "auth/login.html"
	renderTemplateWithData(w, path, nil)
}

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		renderTemplateWithData(w, "auth/register.html", nil)
		return
	}

	db, err := dbConnection()
	if err != nil {
		handleError(w, "Erreur DB", http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	params_reg := extractReg(r)
	err = insertUser(db, params_reg)
	if err != nil {
		handleError(w, "Erreur lors de l'insertion dans la database", http.StatusInternalServerError, err)
		return
	}

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func LogoutBtn(w http.ResponseWriter, r *http.Request) {
	Logout(w, r)
	path := "index.html"
	renderTemplateWithData(w, path, nil)
}

func Home(w http.ResponseWriter, r *http.Request) {

	db, err := dbConnection()
	if err != nil {
		handleError(w, "Erreur DB", http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	if r.Method == http.MethodPost {
		params_log := extractLog(r)
		l, err := logUser(db, params_log)
		if err != nil {
			handleError(w, "Erreur lors de la connexion", http.StatusInternalServerError, err)
			return
		}
		if l {
			InitStartSession(w, r)
			http.Redirect(w, r, "/home", http.StatusSeeOther)
			return
		}
	}

	if err := Authorize(r); err == nil {
		path := "protected/home.html"
		messages, err := getMess(db)
		if err != nil {
			handleError(w, "Erreur dans le chargement des messages", http.StatusInternalServerError, err)
			return
		}
		search := extractQueryParams(r)
		data := DataHome{
			Mess:   Search(messages, search.SearchQuery, search.SearchType),
			Params: search,
		}
		renderTemplateWithData(w, path, data)
		return
	}

	path := "auth/login.html"
	renderTemplateWithData(w, path, nil)
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	var new_data any
	db, err := dbConnection()
	if err != nil {
		handleError(w, "Erreur DB", http.StatusInternalServerError, err)
		return
	}

	if err := Authorize(r); err != nil {
		handleError(w, "Utilisateur non authentifié", http.StatusUnauthorized, err)
		return
	}

	data, err := getMess(db)
	if err != nil {
		handleError(w, "Erreur dans le chargement des messages", http.StatusInternalServerError, err)
		return
	}

	path := "protected/home.html"

	search := extractQueryParams(r)
	new_data = DataHome{
		Mess:   Search(data, search.SearchQuery, search.SearchType),
		Params: search,
	}
	defer db.Close()
	renderTemplateWithData(w, path, new_data)
}

func createForum(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		renderTemplateWithData(w, "protected/create.html", nil)
		return
	}

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
	if err := postMess(db, user_id, params_mess.subject, params_mess.tags, params_mess.body); err != nil {
		handleError(w, "Erreur lors de l'insertion du post", http.StatusInternalServerError, err)
		return
	}

	http.Redirect(w, r, "/HomePage", http.StatusSeeOther)
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
