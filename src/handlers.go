package main

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
)

type Message struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func handler(r http.ResponseWriter,req *http.Request) {
	tasks := fetchData()
	todoList,err := json.Marshal(tasks)
	checkError(err)
	r.Write(todoList)
}

func createHandler(writer http.ResponseWriter, request *http.Request) {
	b,_ := ioutil.ReadAll(request.Body)
	defer request.Body.Close()
	var msg Message
	json.Unmarshal(b, &msg)
	con := getDbConnect()
	con.Exec(getQueries("write"),msg.Title,msg.Description)
	defer con.Close()
	writer.Write(b)
}
