package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"io/ioutil"
)

var addr = flag.String("addr", ":8080", "addr to bind to")

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.RemoteAddr, r.Method, r.URL.String())

	fmt.Fprintf(w, "<H1>Adopter Check API landing page</H1>")
	fmt.Fprintf(w, "<H3>Executing the following brownfield API: ")
	fmt.Fprintf(w, r.URL.String())
	fmt.Fprintf(w, "\n")
	
	url := r.URL.String()
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Fprintf(w, string(body))
	
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
