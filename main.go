package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "<h1>Welcome to awesome site</h1>"); 
}

//main function to tell for server startup
func main(){
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe("3000", nil)
}