package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var port = ":8080"

type todo struct {
	title       string
	isCompleted bool
}

type resp struct {
	todos []todo
}

func getTodosHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	firstTodo := todo{title: "First todo", isCompleted: false}
	secondTodo := todo{title: "Second todo", isCompleted: false}
	recs := &resp{todos: []todo{firstTodo, secondTodo}}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(recs); err != nil {
		w.WriteHeader(500)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	t, _ := template.ParseFiles("./public/templates/index.html")

	t.Execute(w, nil)
}

func main() {
	router := httprouter.New()

	router.GET("/todos", getTodosHandler)
	router.GET("/", indexHandler)
	router.ServeFiles("/public/*filepath", http.Dir("./public"))

	if err := http.ListenAndServe(port, router); err != nil {
		fmt.Printf("Something went wrong %s", err)
	} else {
		fmt.Printf("Server listens on port %s", port)
	}
}
