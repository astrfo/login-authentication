package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("secret-key"))

var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
	"user3": "password3",
}

func main() {
	http.HandleFunc("/", loginHandler)
	http.HandleFunc("/home", homeHandler)
	http.HandleFunc("/logout", logoutHandler)
	http.HandleFunc("/signup", signupHandler)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("templates"))))

	log.Println("Server started on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")

		if pass, ok := users[username]; ok && pass == password {
			session, _ := store.Get(r, "session")
			session.Values["username"] = username
			session.Save(r, w)
			http.Redirect(w, r, "/home", http.StatusFound)
		} else {
			renderTemplate(w, "login.html", "Invalid username or password")
		}
	} else {
		renderTemplate(w, "login.html", nil)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	username, ok := session.Values["username"].(string)
	if !ok || username == "" {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	renderTemplate(w, "index.html", username)
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	session.Values["username"] = ""
	session.Save(r, w)
	http.Redirect(w, r, "/", http.StatusFound)
}

func signupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		username := r.FormValue("username")
		password := r.FormValue("password")

		fmt.Println("New user registration:")
		fmt.Println("Username:", username)
		fmt.Println("Password:", password)

		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		renderTemplate(w, "signup.html", nil)
	}
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	tmpl = fmt.Sprintf("templates/%s", tmpl)
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
