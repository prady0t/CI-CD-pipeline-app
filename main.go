package main

import (
	"fmt"
	"net/http"
)

func root(w http.ResponseWriter, e *http.Request){
	fmt.Fprint(w, "Welcome to Home!")
}

func main(){
	http.HandleFunc("/", root)
	fmt.Println("Serving at port 3000")
	err := http.ListenAndServe(":3000", nil)
	if err != nil{
		fmt.Println("Error Occured!")
	}
	
}