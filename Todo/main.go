package main

import (
	"html/template"
	"log"
	"net/http"
)

var todoList []string

// template engine, blank identifier to ignore error
func handleTodo(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/todo.html")
	if err != nil {
		log.Printf("Template parsing error: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
}

err = t.Execute(w, todoList)
if err != nil {
	log.Printf("Template executing error: %v", err)
	http.Error(w, "Internal server error", http.StatusInternalServerError)
	return
}
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