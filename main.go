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

	fmt.Fprintf(w, "hello dekt\n")
	fmt.Fprintf(w, "url.string:\n")
	fmt.Fprintf(w, r.URL.String())
	fmt.Fprintf(w, "revision:\n")
	fmt.Fprintf(w, os.Getenv("REV"))
}

func main() {
	http.HandleFunc("/", handler)

	log.Printf("listening on %s", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
