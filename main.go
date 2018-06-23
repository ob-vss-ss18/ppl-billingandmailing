package main

import (
	"net/http"
	"os"
	"github.com/ob-vss-ss18/ppl-billingandmailing/api"
	"log"
)

func main() {
	// GET PORT FROM HEROKU ENVIRONMENT VARIABLE. IF NOT EXISTENT USE PORT 8000
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	http.Handle("/graphql", api.Init_graphql())
	log.Println("Server started at http://localhost:" + port + "/graphql")
	err := http.ListenAndServe(":" + port, nil)
	if err != nil {
		log.Fatal(err)
		}
	}
