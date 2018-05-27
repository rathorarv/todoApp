package main

var queries  = make(map[string]string)

func getQueries(s string) string {
queries["fetch"] = "select * from todo"
queries["write"] = `insert into todo(title,description)values($1,$2)`
	return queries[s]
}

