package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreetHandler(w http.ResponseWriter, r *http.Request) {
	Greet(os.Stdout, "world")
}

func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreetHandler)))
}
