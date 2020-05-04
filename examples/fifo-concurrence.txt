package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
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

	usersIdChan := make(chan int)

	go func() {
		for _, u := range usersRepo {
			usersIdChan <- u
		}
		close(usersIdChan)
	}()

	workers := make([]<-chan User, 2)

	workers = append(workers, getUserDetail(usersIdChan))
	workers = append(workers, getUserDetail(usersIdChan))

	endUserChan := make(chan User)

	for _, chanU := range workers {
		go func(chanU <-chan User, endUserChan chan User) {
			for u := range chanU {
				endUserChan <- u
			}
		}(chanU, endUserChan)
	}

	for x := 0; x < len(usersRepo); x++ {
		users = append(users, <-endUserChan)
	}

	sort.Slice(users, func(i, j int) bool {
		return users[i].Id < users[j].Id
	})

	return users
}

func getUserDetail(usersChan chan int) <-chan User {

	userDetailChan := make(chan User)

	for userId := range usersChan {

		go func(userId int) {

			req, _ := http.NewRequest("GET", "http://users-service:8090/user", nil)

			q := req.URL.Query()
			q.Add("userId", strconv.Itoa(userId))
			req.URL.RawQuery = q.Encode()

			client := http.Client{}
			resp, _ := client.Do(req)

			body, _ := ioutil.ReadAll(resp.Body)

			user := User{}
			json.Unmarshal(body, &user)

			userDetailChan <- user
		}(userId)

	}

	return userDetailChan

}
