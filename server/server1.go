package main

// package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func main(){

	fs := http.FileServer(http.Dir("./static"));
	http.Handle("/math-form/", http.StripPrefix("/math-form/", fs));

	http.HandleFunc("/add-form", func(w http.ResponseWriter, r* http.Request){
		if r.Method != http.MethodPost{
			http.Error(w, "Invalid Request", http.StatusMethodNotAllowed);
			return
		}

		if err := r.ParseForm(); err != nil {
			http.Error(w, "Cant parse body", http.StatusBadRequest)
			return
		}

		num1Str := r.FormValue("num1");
		num2Str := r.FormValue("num2");

		num1Int, err1 := strconv.Atoi(num1Str);
		num2Int, err2 := strconv.Atoi(num2Str);

		if err1 != nil || err2 != nil {
			http.Error(w, "Invalid Input", http.StatusBadRequest);
		}

		result := num1Int + num2Int

		response := map[string]interface{}{
			"result" : result,
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response);


	})

	http.HandleFunc("/add", func(w http.ResponseWriter, r* http.Request){
		r.ParseForm();
		a := r.FormValue("a");
		b := r.FormValue("b");

		aInt, err1 := strconv.Atoi(a);
		bInt, err2 := strconv.Atoi(b);

		if err1 != nil || err2 != nil {
			http.Error(w, "Invalid Input", http.StatusBadRequest);
		}

		fmt.Fprintf(w, "Addition of %d and %d is %d", aInt, bInt, aInt+bInt);
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprint(w, "Greetings");
	})

	fmt.Println("Server is running on port: 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server is not running...");
	}
}