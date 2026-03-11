package main

import (
	"fmt";
	"encoding/json"
	"net/http"
)

type User struct{
	ID int `json:"id"`;
	Name string `json:"name"`;
	Email string `json:"email"`;
}

func homeHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path == "/" || r.URL.Path == "/home"{
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, "<h1>Welcome to my webpage</h1><p>Namaste!!!!</p>")
		return
	}
	http.NotFound(w,r);
}

func aboutHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>About us : </h1><p>We are learning golang...</p>");
	// return
}

func userHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-type", "text/html")
	users := []User{
		{ID: 1, Name: "Amith", Email: "amith@redifff.com"},
		{ID: 2, Name: "Arjay", Email: "arjay@hotmail.com"},
		{ID: 1, Name: "Rohan", Email: "gmail@rohan.com"},
	}
	json.NewEncoder(w).Encode(users);
}