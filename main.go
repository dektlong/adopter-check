package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"io/ioutil"
	"strings"
)

var addr = flag.String("addr", ":8080", "addr to bind to")
var apiCall string

func handler(w http.ResponseWriter, r *http.Request) {
	
	log.Println(r.RemoteAddr, r.Method, r.URL.String())
	
	fmt.Fprintf(w, "<H1>")
	fmt.Fprintf(w, "Welcome to the Adopter Check function")
	fmt.Fprintf(w, "</H1>")
	
	fmt.Fprintf(w, "<H3>")
	fmt.Fprintf(w, "Please provide a valid 'brownfield API' for execution")
	
	log.Println(r.RemoteAddr, r.Method, r.URL.String())
	
	apiCall = strings.Replace(r.URL.String(), "api=", "", 1)
	
	fmt.Fprintf(w, "<H3>")
	fmt.Fprintf(w, "Executing the following brownfield API: ")
	fmt.Fprintf(w, apiCall)
	fmt.Fprintf(w, "\n")
	
	req, _ := http.NewRequest("GET", apiCall, nil)
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
	
	http.HandleFunc("/", roothandler)

	log.Printf("listening on %s", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
