package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/form", formHandler)

	port := ":8000"
	server := &http.Server{
		Addr:           port,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	fmt.Printf("Starting server on port %s", port)
	log.Fatal(server.ListenAndServe())
}

func helloHandler(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/hello" {
		http.Error(writer, "404 -- Not found", http.StatusNotFound)
		return
	}
	if request.Method != "GET" {
		http.Error(writer, "method not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(writer, "Hello, Sojourner!")
}

func formHandler(writer http.ResponseWriter, request *http.Request) {
	if err := request.ParseForm(); err != nil {
		fmt.Fprintf(writer, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(writer, "POST request successful")
	name := request.FormValue("name")
	address := request.FormValue("address")

	fmt.Fprintf(writer, "Name = %s\n", name)
	fmt.Fprintf(writer, "Address = %s\n", address)
}
