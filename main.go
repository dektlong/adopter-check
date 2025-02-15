package main

import (
	"flag"
	"io/ioutil"
	"fmt"
	"log"
	"net/http"
)

var API_CALL="http://brownfield.example.com/api/try-me"

func handler(w http.ResponseWriter, r *http.Request) {
	
	log.Println(r.RemoteAddr, r.Method, r.URL.String())
	
	fmt.Fprintf(w, "<H1><font color='navy'>Welcome to Adopter Check function</font></H1>")
        fmt.Fprintf(w, "<H2><font color='gray'>Brownfield API: </font>")
	fmt.Fprintf(w, "<font color='maroon'>")
	fmt.Fprintf(w, API_CALL)
	fmt.Fprintf(w, "</font><BR><BR>")
	fmt.Fprintf(w, "<font color='gray'>Response: </font>")

	response, err := http.Get(API_CALL)

	if err != nil {

    		log.Println(err.Error())
		fmt.Fprintf(w, "<font color='red'>")
		fmt.Fprintf(w,"Adoption Denied")
		fmt.Fprintf(w, "</font>")
		fmt.Fprintf(w, "<BR><BR><img src='https://raw.githubusercontent.com/dektlong/adopter-check/main/sad-dog.jpg' alt=''>")
		
    } else {
		responseData, err := ioutil.ReadAll(response.Body)
		log.Println(string(responseData),err)
		fmt.Fprintf(w,"Adoption Approved. Get ready to meet your new best friend.")
		fmt.Fprintf(w, "</font>")
		fmt.Fprintf(w, "<BR><BR><img src='https://raw.githubusercontent.com/dektlong/adopter-check/main/happy-dog.jpg' alt=''>")
	}
	fmt.Fprintf(w, "</H2>")
}


func main() {
	
	http.HandleFunc("/", handler)

	var addr = flag.String("addr", ":8080", "addr to bind to")
	log.Printf("listening on %s", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
