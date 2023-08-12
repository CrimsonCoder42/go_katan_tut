package main

import (
	"fmt"
	"net/http"
)

func main() {
	//http.HandleFunc("/create-payment-intent", handleCreatePaymentIntent)
	//http.HandleFunc("/health", handleHealth)
	//log.Println("Listening on http://localhost:4242")
	//var err error = http.ListenAndServe("localhost:4242", nil)
	//if err != nil {
	//	log.Fatal("ListenAndServe: ", err)
	//}

	var names []string = []string{"James", "Jill", "Jack"}
	fmt.Println(names[2])
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
}

func handleCreatePaymentIntent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ENDPOINT CALLED")
}
