package main

import (
	"net/http"
)


func main() {
	dbConnection := getDbConnect()
	http.HandleFunc("/",getHandler(dbConnection))
	http.HandleFunc("/create", getCreateHandler(dbConnection))
	err := http.ListenAndServe(":8000",nil)
	if err != nil {
		return
	}
	defer dbConnection.Close()
}
