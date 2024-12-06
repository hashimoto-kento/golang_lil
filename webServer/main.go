// run the server with the command: go run main.go, then open the browser and type: http://localhost:8080/ to see the output
package main

import (
	//"fmt"
	"log"
	"net/http"
)

//func hello(w http.ResponseWriter, r *http.Request) { // create a handler function
//	fmt.Fprintf(w, "Hello World!, and application") // display on the literal
//}

func main() {
	http.Handle("/", http.FileServer(http.Dir("static"))) // set router
	err := http.ListenAndServe(":8080", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}