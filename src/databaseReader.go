package src

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host   = "localhost"
	port   = 5431
	dbName = "arvinds"
)

func GetDbConnect() *sql.DB {
	dbInfo := fmt.Sprintf("host=%s port=%d dbname=%s sslmode=disable",
		host, port, dbName)
	connection, err := sql.Open("postgres", dbInfo)
	CheckError(err)
	return connection
}

type Todo struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func fetchData(connection *sql.DB) []Todo {
	tasks, err := connection.Query(getQueries("fetch"))
	CheckError(err)
	defer tasks.Close()
	return todos(tasks)
}

func todos(rows *sql.Rows) []Todo {
	todoList := make([]Todo, 0)
	for rows.Next() {
		var id int
		var title, description string
		rows.Scan(&id, &title, &description)
		todo := Todo{id, title, description}
		todoList = append(todoList, todo)
	}
	return todoList
}
