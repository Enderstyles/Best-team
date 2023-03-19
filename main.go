package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"unicode"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
)

var db *sql.DB

var store = sessions.NewCookieStore([]byte("secret-key"))

type User struct {
	ID       int
	Fullname string
	Email    string
	Username string
	Password string
}
type Items struct {
	ID      int
	Name    string
	Content string
	Picture string
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
		err := res.Scan(&user.ID, &user.Username, &user.Password, &user.Email, &user.Fullname)

		if err != nil {
			return err
		}
		
	}
	return nil
}

func register(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.ServeFile(w, r, "views/register.html")
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fullName := r.FormValue("fullName")
	email := r.FormValue("email")
	username := r.FormValue("username")
	password := r.FormValue("password")

	// Check for required fields
	if fullName == "" || email == "" || username == "" || password == "" {
		http.Error(w, "Required field is missing", http.StatusBadRequest)
		return
	}

	// Check for correct email format
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		http.Error(w, "Please enter the correct email address", http.StatusBadRequest)
		return
	}

	//Check for full name format
	nameRegex := regexp.MustCompile(`^[a-zA-Z]+\s+[a-zA-Z]+$`)
	if !nameRegex.MatchString(fullName) {
		http.Error(w, "The full user name must contain: 'First name' and 'Last Name' required", http.StatusBadRequest)
		return
	}

	// Check for length of user name
	if len(username) < 5 || len(username) > 30 {
		http.Error(w, "The length of the user name must be from 5 to 30 characters", http.StatusBadRequest)
		return
	}

	if len(password) < 8 {
		http.Error(w, "Password must be at least 8 characters long", http.StatusBadRequest)
		return
	}

	hasNumber := false
	hasUpper := false
	hasLower := false
	for _, c := range password {
		switch {
		case unicode.IsNumber(c):
			hasNumber = true
		case unicode.IsUpper(c):
			hasUpper = true
		case unicode.IsLower(c):
			hasLower = true
		}
	}

	if !(hasNumber && hasUpper && hasLower) {
		http.Error(w, "Password must contain at least one digit, uppercase letter, lowercase letter", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = db.Exec("INSERT INTO users (full_name, email, username, password) VALUES (?, ?, ?, ?)", fullName, email, username, hashedPassword)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Redirect to login page
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
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

	// Check for required fields
	if username == "" || password == "" {
		http.Error(w, "Username and Password are required", http.StatusBadRequest)
		return
	}

	// Retrieve user from database
	var user User
	row := db.QueryRow("SELECT username, password FROM users WHERE username = ?", username)
	err = row.Scan(&user.Username, &user.Password)
	if err == sql.ErrNoRows {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Check password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

// Home page
func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "views/index.html")
}

func searchitems(w http.ResponseWriter, r *http.Request) {
	// Getting the search query from the form
	query := r.FormValue("query")

	// Getting the list of items from search query
	items, err := search(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Execute search template with the list of items
	tmpl, err := template.ParseFiles("templates/search.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, items)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func search(query string) ([]Items, error) {
	// Split the search query into individual words
	words := strings.Split(query, " ")

	// Building the search query
	var where []string
	var args []interface{}
	for _, word := range words {
		if len(word) > 0 {
			where = append(where, "MATCH(name,content) AGAINST(? IN BOOLEAN MODE)")
			args = append(args, word+"*")
		}
	}
	if len(where) == 0 {
		return nil, nil
	}
	whereStr := strings.Join(where, " OR ")
	queryStr := fmt.Sprintf("SELECT id, name, content, picture FROM items WHERE %s", whereStr)

	// Executing the search query
	rows, err := db.Query(queryStr, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Building the list of items
	var items []Items
	for rows.Next() {
		var item Items
		err := rows.Scan(&item.ID, &item.Name, &item.Content, &item.Picture)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	for i, j := 0, len(items)-1; i < j; i, j = i+1, j-1 {
		items[i], items[j] = items[j], items[i]
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return items, nil
}
func allPosts(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id,name, content, picture FROM Items")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var items []Items
	for rows.Next() {
		var item Items
		err = rows.Scan(&item.ID, &item.Name, &item.Content, &item.Picture)

		if err != nil {
			panic(err.Error())
		}
		items = append(items, item)
	}
	for i, j := 0, len(items)-1; i < j; i, j = i+1, j-1 {
		items[i], items[j] = items[j], items[i]
	}
	tmpl, err := template.ParseFiles("templates/search.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, items)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
func createItem(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.ServeFile(w, r, "views/create_item.html")
		return
	}
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//getting values from form
	name := r.FormValue("name")
	content := r.FormValue("content")

	if name == "" || content == "" {
		http.Error(w, "Required field is missing", http.StatusBadRequest)
		return
	}

	//getting image from form
	file, handler, err := r.FormFile("img")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	//creating file on server side to save image
	f, err := os.OpenFile("C:/xampp/htdocs/pictures/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer f.Close()

	//copy the uploaded img to file
	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//inserting img path to db
	file_location := fmt.Sprintf("%s%s", "pictures/", handler.Filename)
	_, err = db.Exec("INSERT INTO Items (name, content, picture) VALUES(?,?,?)", name, content, file_location)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/feed", http.StatusFound)
}

func main() {
	// Connecting to mysql
	err := Connect()
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()

	r.HandleFunc("/register", register).Methods("GET", "POST")
	r.HandleFunc("/login", login).Methods("GET", "POST")

	r.HandleFunc("/", home)

	r.HandleFunc("/search", searchitems)
	r.HandleFunc("/create_item", createItem)
	r.HandleFunc("/feed", allPosts)

	// Serve static files
	fs := http.FileServer(http.Dir("static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	fmt.Println("Server path: http://localhost:3000")
	http.ListenAndServe(":3000", r)
	if err != nil {
		log.Fatal(err)
	}
}
