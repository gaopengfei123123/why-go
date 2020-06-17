package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("run")
	http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))
}
