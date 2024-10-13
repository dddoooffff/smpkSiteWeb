package service

import (
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"
)

type Users struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var User []Users

func MainPage(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method == http.MethodPost {
		name := r.FormValue("name")
		email := r.FormValue("email")

		anFunc := AddUser(name, email)

		n, em, id := anFunc()

		strID := strconv.Itoa(id)

		User = append(User, Users{strID, n, em})
		http.Redirect(w, r, "/", http.StatusSeeOther)
		id++
	}

	if err := temp.Execute(w, "home"); err != nil {
		http.Error(w, "Failed to render tamplate", http.StatusInternalServerError)
	}
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(User)
}
