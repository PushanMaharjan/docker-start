package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "welcome to home page")
	fmt.Printf("here at homepage")
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

//users

func getUsers() []*User {
	db, err := sql.Open("mysql", "test_user:secret@tcp(db:3306)/test_db")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	results, err := db.Query("SELECT * FROM users")
	if err != nil {
		panic(err)
	}

	var users []*User

	for results.Next() {
		var u User

		err = results.Scan(&u.ID, &u.Name)
		if err != nil {
			panic(err)
		}

		users = append(users, &u)
	}
	return users
}

func userPage(w http.ResponseWriter, r *http.Request) {

	users := getUsers()

	fmt.Fprintf(w, "users endpoint")
	json.NewEncoder(w).Encode(users)
}

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/users", userPage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
