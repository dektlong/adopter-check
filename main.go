package main

import (
	"flag"
	"io/ioutil"
	"fmt"
	"log"
	"net/http"
	"os"
)

var API_CALL="http://dekt4pets.tanzu.dekt.io/api/animals"

func handler(w http.ResponseWriter, r *http.Request) {
	
	log.Println(r.RemoteAddr, r.Method, r.URL.String())
	
	fmt.Fprintf(w, "<H1>Welcome to Adopter Check function</H1>")
		    
	fmt.Fprintf(w, "<H2>Running Brownfield API(s) ...</H2>")
	
	fmt.Fprintf(w, "<H3>")

	fmt.Fprintf(w, "API:")
	fmt.Fprintf(w, API_CALL)
	fmt.Fprintf(w, "<BR>")

	response, err := http.Get("http://dekt4pets.tanzu.dekt.io/api/animals")

	if err != nil {
    	log.Println(err.Error())
		fmt.Fprintf(w,"Unable to exectute this API")
    }

    responseData, err := ioutil.ReadAll(response.Body)
	
    if err != nil {
		log.Println(err.Error())
		fmt.Fprintf(w,"Unable to exectute this API")
    } else {
		fmt.Fprintf(w,string(responseData))
	}
	
    	
	
	fmt.Fprintf(w, "</H3>")
	
	fmt.Fprintf(w, "<H4>Function revision: ")
	fmt.Fprintf(w, os.Getenv("REV"))
	fmt.Fprintf(w, "</H4>")
}


func main() {
	
	http.HandleFunc("/", handler)

	var addr = flag.String("addr", ":8080", "addr to bind to")
	log.Printf("listening on %s", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
