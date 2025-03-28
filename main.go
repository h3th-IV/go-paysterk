package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	secretKey := "sk_test_fa16b06664111cf77ebcd2df5d58a1110ca0dfa6"
	url := "https://api.paystack.co/transaction/initialize"

	reqBody := RequestBody{
		Email:  "customer@email.com",
		Amount: 500000,
	}

	jsonReq, err := json.Marshal(reqBody)
	if err != nil {
		fmt.Println(err)
		return
	}
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(jsonReq))
	if err != nil {
		log.Printf("err creating request: %v", err)
		return
	}
	req.Header.Set("Authorization", "Bearer "+secretKey)
	req.Header.Set("Content-Type", "application/json")

	var Client http.Client
	res, err := Client.Do(req)
	if err != nil {
		log.Printf("err sending request: %v", err)
		return
	}
	defer res.Body.Close()
	resByte, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("err reading response body: %v", err)
		return
	}
	fmt.Printf("response: %v\n", string(resByte))
}

type RequestBody struct {
	Email  string `json:"email"`
	Amount int64  `json:"amount"`
}
