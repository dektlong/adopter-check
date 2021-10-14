package main

import (
	"flag"
	"io/ioutil"
	"fmt"
	"log"
	"net/http"
)

var API_CALL="http://datacheck.apps.dekt.io/api/adoption-history?adopterID=99-999-9999"

//https://brownfieldapi.example.com/my-brownfield-api

func handler(w http.ResponseWriter, r *http.Request) {
	
	log.Println(r.RemoteAddr, r.Method, r.URL.String())
	
	fmt.Fprintf(w, "<H1><font color='teal'>")
	fmt.Fprintf(w, "Welcome to Adopter Check function")
	fmt.Fprintf(w, "</font></H1>")
	
	fmt.Fprintf(w, "<H2>")
	fmt.Fprintf(w, "Brownfield API: ")
	fmt.Fprintf(w, "<H1><font color='blue'>")
	fmt.Fprintf(w, API_CALL)
	fmt.Fprintf(w, "</font>")
	fmt.Fprintf(w, "<BR><BR>Response: ")

	response, err := http.Get(API_CALL)

	if err != nil {
    		log.Println(err.Error())
		fmt.Fprintf(w, "<H1><font color='red'>")
		fmt.Fprintf(w,"Adoption Denied")
		fmt.Fprintf(w, "</font>")
		
    } else {
		responseData, err := ioutil.ReadAll(response.Body)
		log.Println(string(responseData),err)
		fmt.Fprintf(w, "<H1><font color='green'>")
		fmt.Fprintf(w,"Adoption Approved. Get ready to meet your new best friend.")
		fmt.Fprintf(w, "</font>")
	}
	fmt.Fprintf(w, "</H2>")
}


func main() {
	
	http.HandleFunc("/", handler)

	var addr = flag.String("addr", ":8080", "addr to bind to")
	log.Printf("listening on %s", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
