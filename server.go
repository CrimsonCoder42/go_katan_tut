package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/create-payment-intent", handleCreatePaymentIntent)
	http.HandleFunc("/health", handleHealth)
	log.Println("Listening on http://localhost:4242")
	var err = http.ListenAndServe("localhost:4242", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

	var names = []string{"James", "Jill", "Jack"}
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(numbers[3])
	fmt.Println(names[0])
}

func handleCreatePaymentIntent(writer http.ResponseWriter, request *http.Request) {

	if request.Method != "POST" {
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		ProductId string `json:"product_id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Address1  string `json:"address_1"`
		Address2  string `json:"address_2"`
		City      string `json:"city"`
		State     string `json:"state"`
		Zip       string `json:"zip"`
		Country   string `json:"country"`
	}

	err := json.NewDecoder(request.Body).Decode(&req)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(req.ProductId)

	params := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(calculatedOderAmount(req.ProductId)),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
	}
}

func handleHealth(writer http.ResponseWriter, request *http.Request) {
	response := []byte("server is up and running!")

	_, err := writer.Write(response)
	if err != nil {
		fmt.Println(err)
	}

}
