package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

var usersRepo = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

func main() {

	http.HandleFunc("/proxy/users/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(getAll())
	})

	fmt.Println("Starting web server.")

	http.ListenAndServe(":8070", nil)
}

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func getAll() []User {
	users := []User{}

	for _, id := range usersRepo {
		req, _ := http.NewRequest("GET", "http://users-service:8090/user", nil)

		q := req.URL.Query()
		q.Add("userId", strconv.Itoa(id))
		req.URL.RawQuery = q.Encode()

		client := http.Client{}
		resp, _ := client.Do(req)

		body, _ := ioutil.ReadAll(resp.Body)

		user := User{}
		json.Unmarshal(body, &user)
		users = append(users, user)
	}

	return users
}
