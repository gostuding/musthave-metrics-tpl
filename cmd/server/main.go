package main

import (
	"net/http"

	"github.com/gostuding/musthave-metrics-tpl/cmd/server/handlers"
)

func main() {
	myServerMux := http.NewServeMux()
	myServerMux.HandleFunc("/", handlers.PathNotFound)
	myServerMux.HandleFunc("/update/", handlers.Update)

	err := http.ListenAndServe(`:8080`, myServerMux)
	if err != nil {
		panic(err)
	}
}
