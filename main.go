package main

import ("fmt"
	"net/http"
	"log"
	"os"
)

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server started...")
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))		//  assigned by Heroku as the PORT environment variable
}

func handler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello World (Dev Branch)")
}
