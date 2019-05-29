package main

import (
	"net/http"

	src "github.com/rathourarv/todoApp/src"
)

func main() {
	dbConnection := src.GetDbConnect()
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)
	http.HandleFunc("/create", src.Create(dbConnection))
	http.HandleFunc("/todo", src.GetHandler(dbConnection))
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		return
	}
	defer dbConnection.Close()
}
