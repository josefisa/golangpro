package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method is not supported", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprint(w, "hello!")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method is not supported!", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() error: %v", err)
		return
	}
	fmt.Fprintln(w, "POST request succesfull")

	name := r.FormValue("name")
	address := r.FormValue("adress")

	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Adress = %s\n", address)
}

func main() {
	fileServer := http.FileServer(http.Dir("."))

	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Starting from server 8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
