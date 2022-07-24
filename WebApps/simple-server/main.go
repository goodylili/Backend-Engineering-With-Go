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
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/", formHandler)



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
	
}

func formHandler(writer http.ResponseWriter, request *http.Request) {
	
}
}