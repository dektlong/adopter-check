package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

var API_CALL="datacheck.tanzu.dekt.io/api/adoption-history?adopterID=909-99-9999"

func handler(w http.ResponseWriter, r *http.Request) {
	
	log.Println(r.RemoteAddr, r.Method, r.URL.String())
	
	fmt.Fprintf(w, "<H1>Welcome to Adopter Check function</H1>")
		    
	fmt.Fprintf(w, "<H2>Brownfield API: ")
	fmt.Fprintf(w, API_CALL)
	fmt.Fprintf(w, "</H2>")
	
	fmt.Fprintf(w, "<H3>Function revision: ")
	fmt.Fprintf(w, os.Getenv("REV"))
	fmt.Fprintf(w, "</H3>")
}

func main() {
	
	http.HandleFunc("/", handler)

	var addr = flag.String("addr", ":8080", "addr to bind to")
	log.Printf("listening on %s", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
