package main

import (
	"fmt"
	"net/http"
)

var metricWriter = NewMemStorage()

func main() {
	myServerMux := http.NewServeMux()
	myServerMux.HandleFunc("/update/", update)

	err := http.ListenAndServe(`:8080`, myServerMux)
	if err != nil {
		panic(err)
	}
}

func update(writer http.ResponseWriter, request *http.Request) {
	if request.Method != http.MethodPost {
		// разрешаем только POST-запросы
		fmt.Printf("Method not allowed: method: '%s', path: '%s'\r\n", request.Method, request.URL.Path)
		writer.WriteHeader(http.StatusMethodNotAllowed)
		writer.Write([]byte("Method not allowed. User POST method instead"))
		return
	}
	// status, err := Memory.Add(request.URL.Path)
	status, err := metricWriter.AddMetric(request.URL.Path)
	writer.WriteHeader(status)
	if err != nil {
		fmt.Println(err)
		// writer.Write([]byte(err.Error()))
	} else {
		fmt.Println(metricWriter.Print())
		// writer.Write([]byte(metricWriter.Print()))
	}
}
