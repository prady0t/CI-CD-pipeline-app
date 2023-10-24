package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func root(w http.ResponseWriter, e *http.Request){
	fmt.Fprint(w, "Welcome to Home!")
}

func login(w http.ResponseWriter, e *http.Request){
	response, err := http.Get("http://localhost:3001")
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// Create a Person struct to unmarshal the JSON data
	var person Person

	// Unmarshal the JSON data into the Person struct
	err = json.Unmarshal(body, &person)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	// Display the parsed JSON data
	fmt.Fprint(w, person)
}

func main(){
	http.HandleFunc("/", root)
	http.HandleFunc("/login", login)
	fmt.Println("Serving at port 3000")
	err := http.ListenAndServe(":3000", nil)
	if err != nil{
		fmt.Println("Error Occured!")
	}
	
}