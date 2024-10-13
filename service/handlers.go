package service

import (
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Users struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var User []Users

func MainPage(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("home.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	name := r.FormValue("name")
	email := r.FormValue("email")

	anFunc := AddUser(name, email)

	n, em, id := anFunc()

	idStr := strconv.Itoa(id)
	isEmpty, err := EmptyValues(name, email)

	if !isEmpty {
		http.Error(w, err.Error(), http.StatusVariantAlsoNegotiates)
		return
	}
	User = append(User, Users{idStr, n, em})

	if err := temp.Execute(w, "home"); err != nil {
		http.Error(w, "Failed to render tamplate", http.StatusInternalServerError)
	}
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(User)
}

func GetUserById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	path := mux.Vars(r)

	for _, item := range User {
		if item.ID == path["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(nil)
}
