package main

import ("fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", hello)
}

func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Hello World")
}
