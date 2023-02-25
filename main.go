package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"regexp"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

var db *sql.DB
var err error

type User struct {
	ID       int
	Fullname string
	Email    string
	Username string
	Password string
}
type Post struct {
	ID      int
	Title   string
	Content string
}

func Connect() error {
	var err error
	// Connecting to mysql db
	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/golang_project")
	if err != nil {
		fmt.Println(err)
	}
	// Getting all users
	res, err := db.Query("SELECT * FROM users")

	if err != nil {
		return err
	}

	// Printing information of all users
	for res.Next() {
		var user User
		err := res.Scan(&user.Username, &user.Password, &user.Email, &user.Fullname)

		if err != nil {
			return err
		}
		fmt.Println("%\n", user)
	}
	return nil
}

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

	// Check for required fields
	if full_name == "" || email == "" || username == "" || password == "" {
		http.Error(w, "Required field is missing", http.StatusBadRequest)
		return
	}

	// Check for correct email format
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		http.Error(w, "Please enter the correct email address", http.StatusBadRequest)
		return
	}

	// Check for full name format
	nameRegex := regexp.MustCompile(`^[a-zA-Z] + [a-zA-Z]+$`)
	if !nameRegex.MatchString(full_name) {
		http.Error(w, "The full user name must contain: 'First name' and 'Last Name' required", http.StatusBadRequest)
		return
	}

	// Check for length of user name
	if len(username) < 5 || len(username) > 30 {
		http.Error(w, "The length of the user name must be from 5 to 30 characters", http.StatusBadRequest)
		return
	}

	// Check for length of password and character requirements
	passRegex := regexp.MustCompile(`^(?=.*[A-Z])(?=.*[0-9])(?=.*[a-z]).{8,30}$`)
	if !passRegex.MatchString(password) {
		http.Error(w, "The password length must be from 8 to 30 characters and include one capital letter, one character, and one digit", http.StatusBadRequest)
		return
	}

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

// Home page
func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "views/index.html")
}

func searchPosts(w http.ResponseWriter, r *http.Request) {
	// Getting the search query from the form
	query := r.FormValue("query")

	// Getting the list of posts from search query
	posts, err := search(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Execute search template with the list of posts
	tmpl, err := template.ParseFiles("templates/search.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, posts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func search(query string) ([]Post, error) {
	// Split the search query into individual words
	words := strings.Split(query, " ")

	// Building the search query
	var where []string
	var args []interface{}
	for _, word := range words {
		if len(word) > 0 {
			where = append(where, "MATCH(title,content) AGAINST(? IN BOOLEAN MODE)")
			args = append(args, word+"*")
		}
	}
	if len(where) == 0 {
		return nil, nil
	}
	whereStr := strings.Join(where, " OR ")
	queryStr := fmt.Sprintf("SELECT id, title, content FROM posts WHERE %s", whereStr)

	// Executing the search query
	rows, err := db.Query(queryStr, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Building the list of posts
	var posts []Post
	for rows.Next() {
		var post Post
		err := rows.Scan(&post.ID, &post.Title, &post.Content)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return posts, nil
}
func main() {
	Connect()

	r := mux.NewRouter()

	r.HandleFunc("/search", searchPosts)
	r.HandleFunc("/register", register)
	r.HandleFunc("/login", login)
	r.HandleFunc("/", home)

	fmt.Print("Server path: http://localhost:3000")
	http.ListenAndServe(":3000", r)
}
