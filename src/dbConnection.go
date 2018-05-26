package main

import "database/sql"


func fetchData() string{
	connection, _ := sql.Open("postgres","host=localhost port=5432 dbname=arvinds sslmode=disable")
	defer connection.Close()
	todos, err := connection.Query(getQueries("fetch"))
	checkError(err)
	defer todos.Close()
	var data string
	for todos.Next() {
		var title string
		var description string
		todos.Scan(&title, &description)
		data += title + "  "  + description + "\n"
	}
	return data
}