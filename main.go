package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

var addr = flag.String("addr", ":8080", "addr to bind to")

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.RemoteAddr, r.Method, r.URL.String())

	fmt.Fprintf(w, "<H1>Adopter Check API landing page</H1>")
	fmt.Fprintf(w, "<H2>Executing the following brownfield API: ")
	fmt.Fprintf(w, r.URL.String())
	fmt.Fprintf(w, "</H2>")
	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "<H3>Function revision:")
	fmt.Fprintf(w, os.Getenv("REV"))
	fmt.Fprintf(w, "</H3>")
}

func main() {
	http.HandleFunc("/", handler)

	log.Printf("listening on %s", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
