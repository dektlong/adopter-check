package main

import (
	"flag"
	"fmt"
	"os"
	"net/http"
)

var addr = flag.String("addr", ":8080", "addr to bind to")

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<H2>")
	
	fmt.Fprintf(w, r.URL.String())

	fmt.Fprintf(w, "dekt is running adopter-check\n")
	
	fmt.Fprintf(w, "</H2>")
	
	fmt.Fprintf("My name is: ", os.Getenv("REV"))
}

func main() {
	http.HandleFunc("/", handler)

	log.Printf("listening on %s", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
