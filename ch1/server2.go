package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

// минимальный echo-сервер со счетчиком запросов

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler2)
	http.HandleFunc("/count", counter2)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func counter2(writer http.ResponseWriter, request *http.Request) {
	mu.Lock()
	fmt.Fprintf(writer, "Count = %d\n", count)
	mu.Unlock()
}

func handler2(writer http.ResponseWriter, request *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(writer, "URL.Path = %q\n", request.URL.Path)
}
