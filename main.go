package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
    a "github.com/Enderstyles/Best-team/database"
	_ "github.com/go-sql-driver/mysql"
	//"golang.org/x/crypto/bcrypt"
)


type User struct {
	Id 			int
	Fullname 	string
	Email 		string
    Username 	string
    Password 	string
}
const (  
    username = "root"
    password = ""
    hostname = "127.0.0.1:3306"
    dbname   = "golang_project"
)
func dsn(dbName string) string {  
    return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbName)
}
func Connect(){
	db, err := sql.Open("mysql", dsn(dbname))
    if err != nil {
        panic(err.Error())
    }
    res, err := db.Query("SELECT * FROM users")    

    defer db.Close()

    if err != nil {
        log.Fatal(err)
    }

    for res.Next(){
        var user User
        err := res.Scan(&user.Id, &user.Email, &user.Fullname, &user.Username, &user.Password)

        if err != nil{
            log.Fatal(err)
        }
        fmt.Print("%\n",user)
    }
}

func main() {
    
    http.HandleFunc("/", a.Index)
	
	http.HandleFunc("/register", a.Register)

	fmt.Println("Server jalan di: http://localhost:3000")
	http.ListenAndServe(":3000", nil)
    

}