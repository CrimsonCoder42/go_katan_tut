package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/create-payment-intent", handleCreatePaymentIntent)
	http.HandleFunc("/health", handleHealth)
	log.Println("Listening on http://localhost:4242")
	var err error = http.ListenAndServe("localhost:4242", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

	var names []string = []string{"James", "Jill", "Jack"}
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(numbers[3])
	fmt.Println(names[0])
}

func handleCreatePaymentIntent(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ENDPOINT CALLED")
}

func handleHealth(writer http.ResponseWriter, r *http.Request) {
	response := []byte("server is up and running!")

	_, err := writer.Write(response)
	if err != nil {
		fmt.Println(err)
	}

}
