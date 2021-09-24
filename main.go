package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

var addr = flag.String("addr", ":8080", "addr to bind to")

var API_CALL="datacheck.tanzu.dekt.io/api/adoption-history?adopterID=609-99-9999"

func handler(w http.ResponseWriter, r *http.Request) {
	
	log.Println(r.RemoteAddr, r.Method, r.URL.String())
	
	fmt.Fprintf(w, "<H1>")
	fmt.Fprintf(w, "Welcome to the Adopter Check function")
	fmt.Fprintf(w, "</H1>")
	
	fmt.Fprintf(w, "<H3>")
	fmt.Fprintf(w, "Executing the following brownfield API: ")
	fmt.Fprintf(w, API_CALL)
	fmt.Fprintf(w, "\n")
	
	fmt.Fprintf(w, "</H3>")
	
	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "\n")
	fmt.Fprintf(w, "<H4>")
	fmt.Fprintf(w, os.Getenv("REV"))
	fmt.Fprintf(w, "</H4>")
}

func main() {
	
	http.HandleFunc("/", handler)

	log.Printf("listening on %s", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
