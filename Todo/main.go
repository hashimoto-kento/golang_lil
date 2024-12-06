package main

import (
	"html/template"
	"log"
	"net/http"
)

var todoList []string

// template engine, blank identifier to ignore error
func handleTodo(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/todo.html")
	t.Execute(w, todoList)
}

func main() {
	todoList = append(todoList, "Buy Milk", "Buy Eggs", "Buy Bread")

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/todo", handleTodo)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}