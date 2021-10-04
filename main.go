package main

import (
	"flag"
	"io/ioutil"
	"fmt"
	"log"
	"net/http"
	"os"
)

var API_CALL="you"

func handler(w http.ResponseWriter, r *http.Request) {
	
	log.Println(r.RemoteAddr, r.Method, r.URL.String())
	
	fmt.Fprintf(w, "<H1>Welcome to Adopter Check function</H1>")
		    
	fmt.Fprintf(w, "<H2>")

	fmt.Fprintf(w, "Brownfield API: ")
	fmt.Fprintf(w, API_CALL)
	fmt.Fprintf(w, "<BR><BR>Response: ")

	response, err := http.Get(API_CALL)

	if err != nil {
    	log.Println(err.Error())
		fmt.Fprintf(w,"Unable to run this API")
    } else {
		responseData, err := ioutil.ReadAll(response.Body)
		log.Println(string(responseData),err)
		fmt.Fprintf(w,"Success")
	}
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
