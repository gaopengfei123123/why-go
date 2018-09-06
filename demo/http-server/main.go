package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("route: http://localhost:9090/")
	fmt.Println("route: http://localhost:9090/hello")

	mux := http.NewServeMux()

	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/", pageHandler)

	http.ListenAndServe(":9090", mux)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world"))
}

func pageHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hi baby"))
}
