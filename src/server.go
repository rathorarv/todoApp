package main

import (
	_ "github.com/lib/pq"
	"net/http"
	"fmt"
)

func handler(r http.ResponseWriter,req *http.Request) {
	todos := fetchData()
	fmt.Fprint(r,todos)
}


func main() {
	http.HandleFunc("/",handler)
	err := http.ListenAndServe(":8000",nil)
	if err != nil {
		return
	}
}