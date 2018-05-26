package main

var queries  = make(map[string]string)

func getQueries(s string) string {
queries["fetch"] = "select * from todo"
	return queries[s]
}

