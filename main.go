package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	http.HandleFunc("/name", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Lakshmy !")
	})

	http.HandleFunc("/department", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, csbs!")
	})

	http.HandleFunc("/roll", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, I'm 7376212cb126!")
	})

	http.HandleFunc("/color", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "red")
	})

	fmt.Printf("Server running (port=8081), routes:\n")
	fmt.Printf("http://localhost:8081/hello\n")
	fmt.Printf("http://localhost:8081/name\n")
	fmt.Printf("http://localhost:8081/roll\n")
	fmt.Printf("http://localhost:8081/department\n")
	fmt.Printf("http://localhost:8081/color\n")

	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}
}