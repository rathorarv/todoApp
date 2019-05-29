package src

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Message struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func GetHandler(db *sql.DB) func(http.ResponseWriter, *http.Request) {
	return func(r http.ResponseWriter, req *http.Request) {
		tasks := fetchData(db)
		todoList, err := json.Marshal(tasks)
		if err != nil {
			return
		}
		r.Write(todoList)
	}
}

func Create(db *sql.DB) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		b, _ := ioutil.ReadAll(request.Body)
		defer request.Body.Close()
		var msg Message
		json.Unmarshal(b, &msg)
		db.Exec(getQueries("write"), msg.Title, msg.Description)
		writer.Write(b)
	}
}
