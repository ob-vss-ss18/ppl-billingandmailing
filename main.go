package main

import ("fmt"
	"net/http"
	"log"
	"os"
	"ppl-billingandmailing/pdf"
)

func main() {
	pdf.InvoiceExample()

	http.HandleFunc("/", handler)
	fmt.Println("Server started...")
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), nil))		//  assigned by Heroku as the PORT environment variable
}

func handler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello World (Dev Branch)")
}
