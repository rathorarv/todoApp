package main

import (
	"net/http"

)


func main() {
	dbConnection := getDbConnect()
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/",fs)
	http.HandleFunc("/create", create(dbConnection))
	http.HandleFunc("/todo",getHandler(dbConnection))
	err := http.ListenAndServe(":9000",nil)
	if err != nil {
		return
	}
	defer dbConnection.Close()
}
