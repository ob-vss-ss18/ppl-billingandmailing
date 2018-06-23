package main

import (
	"net/http"
	"os"

	"log"
	"fmt"
	"github.com/ob-vss-ss18/ppl-billingandmailing/api"
	"github.com/ob-vss-ss18/ppl-billingandmailing/pdf"
)

func main() {
	// GET PORT FROM HEROKU ENVIRONMENT VARIABLE. IF NOT EXISTENT USE PORT 8000
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	
	http.HandleFunc("/", handler)
	http.Handle("/graphql", api.Init_graphql())
	log.Println("Server started at http://localhost:" + port + "/graphql")
	err := http.ListenAndServe(":" + port, nil)
	if err != nil {
		log.Fatal(err)
		}
  	pdf.InvoiceExample()
	}

func handler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello World (Dev Branch)")
}
