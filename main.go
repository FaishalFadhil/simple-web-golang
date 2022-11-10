package main

import (
	"encoding/json"
	"errors"
	"html/template"
	"net/http"
)

type List struct {
	ID       int
	Name     string
	Email    string
	Password string
	Address  string
	Job      string
	Reason   string
}

var listPeople = []List{
	{ID: 1, Name: "Joshua", Email: "Joshua@mail.com", Password: "Joshua123", Address: "Bandung", Job: "Programmer", Reason: "Joshua Reason"},
	{ID: 2, Name: "Udin", Email: "Udin@mail.com", Password: "Udin123", Address: "Bali", Job: "Developer", Reason: "Udin Reason"},
	{ID: 3, Name: "Niki", Email: "Niki@mail.com", Password: "Niki123", Address: "New York", Job: "Singer", Reason: "Niki Reason"},
	{ID: 4, Name: "Joji", Email: "Joji@mail.com", Password: "Joji123", Address: "Garut", Job: "Singer", Reason: "Joji Reason"},
}

var PORT = ":8080"

func main() {
	http.HandleFunc("/", greet)
	http.HandleFunc("/list", getData)
	http.HandleFunc("/login", login)

	http.ListenAndServe(PORT, nil)
}

func greet(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("index.html")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	tpl.Execute(w, "Hello, Login")
	return
}

func getData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		json.NewEncoder(w).Encode(listPeople)
		return
	}

	http.Error(w, "invalid method", http.StatusBadRequest)
}

func login(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		email := r.FormValue("email")
		password := r.FormValue("password")

		data, err := findEmail(email, password)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		tpl, errTemp := template.ParseFiles("data.html")

		if errTemp != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		tpl.Execute(w, data)

		return
	}

	http.Error(w, "invalid method", http.StatusBadRequest)
}

func findEmail(email string, password string) (List, error) {
	var data List
	for _, v := range listPeople {
		if v.Email == email {
			if v.Password == password {
				return v, nil
			}
		}
	}
	return data, errors.New("wrong password/email")
}
