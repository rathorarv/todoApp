package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"fmt"
)

const (
	host = "localhost"
	port = 5432
	dbname = "arvinds"
)


func getDbConnect() *sql.DB {
	dbinfo := fmt.Sprintf("host=%s port=%d dbname=%s sslmode=disable",
		host,port,dbname)
	connection, err := sql.Open("postgres",dbinfo)
	checkError(err)
	return connection
}


type Todo struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
}


func fetchData() []Todo{
	connection := getDbConnect()
	tasks, err := connection.Query(getQueries("fetch"))
	checkError(err)
	defer connection.Close()
	defer tasks.Close()
	return convertToList(tasks)
}

func convertToList(rows *sql.Rows) []Todo {
	todoList := make([]Todo,0)
	for rows.Next() {
		var id int
		var title,description string
		rows.Scan(&id ,&title, &description)
		todo := Todo{id ,title,description}
		todoList = append(todoList,todo)
	}
	return todoList
}