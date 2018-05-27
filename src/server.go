package main

import (
	"net/http"
)


func main() {
	http.HandleFunc("/",handler)
	http.HandleFunc("/create", createHandler)
	err := http.ListenAndServe(":8000",nil)
	if err != nil {
		return
	}
}
