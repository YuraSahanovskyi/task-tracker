package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Hello from task service"))

}

func main() {
	fmt.Println("Hello from task service")
	http.HandleFunc("GET /", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
