package controllers

import (
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("views/index.html")

	if err != nil {
		panic(err)
	}
	temp.Execute(w, nil)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/register", http.StatusSeeOther)
}

func Register(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		temp, _ := template.ParseFiles("views/register.html")
		temp.Execute(w, nil)

	}
}