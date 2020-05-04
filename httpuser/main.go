package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

func main() {

	http.HandleFunc("/user/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		value, _ := strconv.Atoi(r.FormValue("userId"))

		time.Sleep(time.Millisecond * 195)
		json.NewEncoder(w).Encode(getUserByID(value))
	})

	http.ListenAndServe(":8090", nil)
}

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func getUsers() []User {
	return []User{
		User{1, "Joao"},
		User{2, "Maria"},
		User{3, "Claudio"},
		User{4, "Fernanda"},
		User{5, "Clara"},
		User{6, "Bianca"},
		User{7, "Mariana"},
		User{8, "Claudio"},
		User{9, "Joana"},
		User{10, "Alexandre"},
	}
}

func getUserByID(userID int) User {
	for _, user := range getUsers() {
		if user.Id == userID {
			return user
		}
	}

	return User{0, "no user"}
}
