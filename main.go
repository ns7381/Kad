package main

import (
	"log"
	"net/http"
	"github.com/ns7381/Kad/handler"
	"github.com/ns7381/Kad/client"
)

func main() {
	clientManager := client.NewClientManager()
	apiHandler, err := handler.CreateHttpAPIHandler(clientManager)
	if err != nil {
		log.Fatalf("create http api handler error.", err)
	}
	http.Handle("/api/", apiHandler)

	log.Printf("start listening on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
