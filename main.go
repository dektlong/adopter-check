package main

import (
	"flag"
//	"io/ioutil"
	"fmt"
	"log"
	"net/http"
	"os"
)

var API_CALL="NONE"

func handler(w http.ResponseWriter, r *http.Request) {
	
	log.Println(r.RemoteAddr, r.Method, r.URL.String())
	
	fmt.Fprintf(w, "<H1>Welcome to Adopter Check function</H1>")
		    
	fmt.Fprintf(w, "<H2>Running Brownfield API(s):</H2>")
	
	fmt.Fprintf(w, "<H3>")

	fmt.Fprintf(w, "API:")
	fmt.Fprintf(w, API_CALL)
	fmt.Fprintf(w, "\nRespose:\n")

	//response, err := http.Get(API_CALL)

//    	if err != nil {
  //      	log.Println(err.Error())
    //    	os.Exit(1)
    //	}

    //responseData, err := ioutil.ReadAll(response.Body)
	
    //	if err != nil {
      //  	log.Fatal(err)
		//	os.Exit(1)
    	//}
    	
//	fmt.Fprintf(w,string(responseData))
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
