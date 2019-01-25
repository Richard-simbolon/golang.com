package main

import (
	"fmt"
	"net/http"
)

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1> WELCOME TO MY JUNGLE !</h1>")
}
func main() {

	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe(":3000", nil)
}
