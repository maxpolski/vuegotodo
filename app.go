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
	Title       string `json:"title"`
	IsCompleted bool   `json:"isCompleted"`
}

var todos = []todo{
	todo{Title: "First Todo", IsCompleted: true},
	todo{Title: "Second Todo", IsCompleted: false},
}

func getTodosHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(todos); err != nil {
		w.WriteHeader(500)
	}
}

func addTodoHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Print(r.Body)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if err := json.NewEncoder(w).Encode(todos); err != nil {
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
	// router.PUT("/todos", addTodoHandler)
	router.GET("/", indexHandler)
	router.ServeFiles("/public/*filepath", http.Dir("./public"))

	if err := http.ListenAndServe(port, router); err != nil {
		fmt.Printf("Something went wrong %s", err)
	} else {
		fmt.Printf("Server listens on port %s", port)
	}
}
