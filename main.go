package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type User struct {
	ID       int
	Fullname string
	Email    string
	Username string
	Password string
}

func Connect() error {
	var err error
	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/golang_project")
	if err != nil {
		fmt.Println(err)
	}

	res, err := db.Query("SELECT * FROM users")

	if err != nil {
		return err
	}

	for res.Next() {
		var user User
		err := res.Scan(&user.Username, &user.Password)

		if err != nil {
			return err
		}
		fmt.Println("%\n", user)
	}
	return nil
}

var db *sql.DB
var err error

func register(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.ServeFile(w, r, "views/register.html")
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	full_name := r.FormValue("full_name")
	email := r.FormValue("email")
	username := r.FormValue("username")
	password := r.FormValue("password")

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	_, err = db.Exec("INSERT INTO users (full_name, email, username, password) VALUES (?, ?, ?, ?)", full_name, email, username, hashedPassword)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Redirect to login page
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.ServeFile(w, r, "views/login.html")
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	username := r.FormValue("username")
	password := r.FormValue("password")

	// Retrieve user from database
	var user User

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
		return
	}
	row := db.QueryRow("SELECT id, username, password FROM users WHERE username = ?", username)
	err = row.Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// Check password
	if user.Password != password {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

func home(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "views/index.html")
}

func main() {
	Connect()

	r := mux.NewRouter()

	// Register routes
	r.HandleFunc("/register", register)
	r.HandleFunc("/login", login)
	r.HandleFunc("/", home)

	fmt.Print("Server path: http://localhost:3000")
	http.ListenAndServe(":3000", r)
}
