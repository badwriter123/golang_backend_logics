// package main
package main

import (
	"fmt"
	"log"
	"net/http"
	
	
)

func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello, I love you %s!", r.URL.Path[1:])
}

func main(){
	http.HandleFunc("/", handler);
	fmt.Println("Server running on port: 8080")
	log.Fatal(http.ListenAndServe(":8080",nil));
}