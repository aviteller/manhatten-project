package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

func signup(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var u user
	err := decoder.Decode(&u)

	if err != nil {
		panic(err)
	}

	if u.Name == "" || u.Email == "" || u.Password == "" {
		http.Error(w, "PLEASE ENTER ALL FIELDS", http.StatusBadRequest)
		return
	}

	fmt.Fprintln(w, u)
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func router(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	path := r.URL.Path
	path = strings.Replace(path, "/", "", -1)

	switch method {
	case "GET":
		switch path {
		case "greet":
			greet(w, r)
		default:
			fmt.Fprintf(w, "GET METHOD: %s DOES NOT EXIST", path)
		}

	case "POST":
		switch path {
		case "signup":
			signup(w, r)

		default:
			fmt.Fprintf(w, "POST METHOD: %s DOES NOT EXIST", path)
		}

	default:
		fmt.Fprintf(w, "METHOD: %s IS NOT ALLOWED", method)
	}
}

func main() {
	http.HandleFunc("/", router)
	http.ListenAndServe(":8080", nil)
}
