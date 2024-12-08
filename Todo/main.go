package main

import (
	"html/template"
	"net/http"
)

var todoList []string

// template engine, blank identifier to ignore error
func handleTodo(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/todo.html")
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
}

err = t.Execute(w, todoList)
if err != nil {
	http.Error(w, "Internal server error", http.StatusInternalServerError)
	return
}
}

func handleAdd(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		todo := r.Form.Get("todo")
		todoList = append(todoList, todo)
		http.Redirect(w, r, "/todo", 303)
}

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/todo", handleTodo)

	http.HandleFunc("/add", handleAdd)

	http.ListenAndServe(":8080", nil)

}