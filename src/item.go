package main

type Item struct {
	item string
}

func (todoItem *Item)toString() string {
	return todoItem.item
}
