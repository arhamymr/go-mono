package main

import (
	"encoding/json"
	"fmt"
	"go-mono/toolkit"
	"log"
	"net/http"
)

const PORT = ":3000"

type welcome string

type User struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type ResponseUser struct {
	status  uint
	data    User
	message string
}

func (wc welcome) Welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello welcome to our server "+string(wc))
}

func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "login")
}

func getJSON(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	switch r.Method {
	case "GET":
		w.Write([]byte(`{
			"data": "hello",
			"message": "get declaraed",
			"matcher": "yuhu"
			}
			`))
	case "POST":
		w.Write([]byte(`"message": "post declaraed"`))
	}
}

func checkUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	var user User
	dbPassword := "arham"

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		log.Fatal("Error decoding into struct", err)
	}

	if user.Password == dbPassword {
		fmt.Println("succesfully login")
	}

	ResUser := &ResponseUser{
		status:  200,
		data:    user,
		message: "login successfully",
	}

	fmt.Println(ResUser, "res user")
	json.NewEncoder(w).Encode(*ResUser)
}

func main() {
	r := http.NewServeMux()

	toolkit.StaticServe("static/")

	var we welcome
	we = "hellow"

	r.HandleFunc("/login", Login)
	r.HandleFunc("/auth/register", checkUser)
	r.HandleFunc("/", we.Welcome)
	r.HandleFunc("/json", getJSON)
	http.ListenAndServe(PORT, r)
}
