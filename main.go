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
	"unicode"

	_ "github.com/go-sql-driver/mysql"
	//"github.com/gorilla/context"
	"github.com/gorilla/mux"

	//"github.com/gorilla/securecookie"
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
	Price   string
}

func Connect() error {
	var err error
	// Connecting to mysql db
	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/golang_project")
	if err != nil {
		fmt.Println(err)
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

func isLoggedIn(r *http.Request) bool {
	// check if the user is logged in by checking if the session contains a user ID
	session, err := store.Get(r, "session")
	if err != nil {
		// handle the error
		return false
	}
	userID, ok := session.Values["userID"].(int)
	fmt.Println("\nIN isLoggedIn :\nsession:  ", session)
	if !ok || userID == 0 {
		fmt.Println("\nIN isLoggedIn:\n", ok, userID)
		return false
	}
	// the user is logged in
	return true
}

func requireLogin(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("IN requireLogin:")
		fmt.Println("\tURL:", r.URL.Path)
		fmt.Println("\tLogged in:", isLoggedIn(r))
		// check if user is logged in
		if !isLoggedIn(r) {
			// redirect to login page
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		// call the next handler function
		next(w, r)
	}
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

	row := db.QueryRow("SELECT id, username, password FROM users WHERE username = ?", username)
	err = row.Scan(&user.ID, &user.Username, &user.Password)
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

	// Add the user ID to the session
	session, err := store.Get(r, "session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	session.Values["authenticated"] = true
	session.Values["userID"] = user.ID
	fmt.Println("\nIn login:\n\tUSER ID:", user.ID)
	err = session.Save(r, w)
	fmt.Println("In login:\n session: ", session)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Redirect the user to the profile page
	http.Redirect(w, r, "/profile", http.StatusSeeOther)
}

func profile(w http.ResponseWriter, r *http.Request) {
	session, err := store.Get(r, "session")
	if err != nil {
		fmt.Println("session error")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Check if the user is authenticated
	if _, ok := session.Values["authenticated"]; !ok {
		fmt.Println("CHECK USER ERROR")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// Get the user's ID from the session
	userID := session.Values["userID"].(int)

	// Get the user's information from the database
	user := User{}
	err = db.QueryRow("SELECT id, full_name, email, username FROM users WHERE id = ?", userID).Scan(&user.ID, &user.Fullname, &user.Email, &user.Username)
	if err != nil {
		fmt.Println("GET USER ERROR : ", userID)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Display the user's information on the profile page
	t, err := template.ParseFiles("views/profile.html")
	if err != nil {
		fmt.Println("template error")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, user)
}

func logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	session.Options.MaxAge = -1
	err := session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// Home page
func home(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("views/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	session, err := store.Get(r, "session")
	if err != nil {
		fmt.Println("session error")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("In home:\nsession: ", session)
	tpl.Execute(w, nil)
}

func createItem(w http.ResponseWriter, r *http.Request) {
	fmt.Println("URL:", r.URL.Path)

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
	price := r.FormValue("price")
	tags := r.FormValue("tags")

	if name == "" || content == "" || price == "" || tags == "" {
		http.Error(w, "Required field is missing", http.StatusBadRequest)
		return
	}

	//getting image from form
	file, handler, err := r.FormFile("img")
	if err != nil {
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
		fileLocation := fmt.Sprintf("%s%s", "pictures/", handler.Filename)
		_, err = db.Exec("INSERT INTO Items (name, content, picture, price, tags) VALUES(?,?,?,?,?)", name, content, fileLocation, price, tags)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	defer file.Close()
	_, err = db.Exec("INSERT INTO Items (name, content, price, tags) VALUES(?,?,?,?)", name, content, price, tags)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/feed", http.StatusFound)
}

func searchitems(w http.ResponseWriter, r *http.Request) {
	// Getting the search query from the form
	query := r.FormValue("query")

	rows, err := db.Query("SELECT id, name, content, picture, price FROM items WHERE name LIKE ?", "%"+query+"%")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	items, err := getItems(rows)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = executeTemplate(w, "templates/search.html", items)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func allItems(w http.ResponseWriter, r *http.Request) {

	rows, err := db.Query("SELECT id, name, content, picture, price FROM Items")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	items, err := getItems(rows)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	reverseItems(items)

	err = executeTemplate(w, "templates/search.html", items)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func minmax(w http.ResponseWriter, r *http.Request) {
	min := r.FormValue("min")
	max := r.FormValue("max")

	rows, err := db.Query("SELECT id, name, content, picture, price FROM items WHERE price BETWEEN ? AND ?", min, max)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	items, err := getItems(rows)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = executeTemplate(w, "templates/search.html", items)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func getItems(rows *sql.Rows) ([]Items, error) {
	var items []Items
	for rows.Next() {
		var item Items
		if err := rows.Scan(&item.ID, &item.Name, &item.Content, &item.Picture, &item.Price); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

func reverseItems(items []Items) {
	for i, j := 0, len(items)-1; i < j; i, j = i+1, j-1 {
		items[i], items[j] = items[j], items[i]
	}
}

func executeTemplate(w http.ResponseWriter, tmpl string, data interface{}) error {
	t, err := template.ParseFiles(tmpl)
	if err != nil {
		return err
	}
	err = t.Execute(w, data)
	if err != nil {
		return err
	}
	return nil
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
	r.HandleFunc("/profile", profile)
	r.HandleFunc("/logout", logout)

	r.HandleFunc("/", home)

	r.HandleFunc("/filter", minmax)
	r.HandleFunc("/search", requireLogin(searchitems))
	r.HandleFunc("/create_item", requireLogin(createItem))
	r.HandleFunc("/feed", requireLogin(allItems))

	// Serve static files
	fs := http.FileServer(http.Dir("static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	fmt.Println("Server path: http://localhost:3000")
	http.ListenAndServe(":3000", r)
	if err != nil {
		log.Fatal(err)
	}
}
