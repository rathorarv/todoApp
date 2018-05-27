package main

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
	"database/sql"
)

type Message struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func getHandler(db *sql.DB) func(http.ResponseWriter,*http.Request){
	return func (r http.ResponseWriter,req *http.Request) {
		tasks := fetchData(db)
		todoList,err := json.Marshal(tasks)
		checkError(err)
		r.Write(todoList)
	}
}

func getCreateHandler(db *sql.DB) func(http.ResponseWriter,*http.Request)  {
	return func (writer http.ResponseWriter, request *http.Request) {
		b,_ := ioutil.ReadAll(request.Body)
		defer request.Body.Close()
		var msg Message
		json.Unmarshal(b, &msg)
		db.Exec(getQueries("write"),msg.Title,msg.Description)
		writer.Write(b)
	}
}
