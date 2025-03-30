## go-paysterk
A Golang library for integrating with Paystack's payment API.


## Overview
go-paysterk is a Go client library for interacting with Paystack's API. It simplifies payment processing, transaction verification, transfers, and more.

## Features
- [x] Initialize and verify transactions
- [x] Manage subscriptions, plans and recurring billing
- [x] Handle bank transfers and mobile money payments
- [ ] Split payments and manage subaccounts
- [ ] Process refunds and chargebacks
- [x] Webhook support for real-time transaction updates
- [x] Charge Usage if you don't want to redirect users to Paystack checkout

## Installation
```sh
go get github.com/h3th-IV/go-paysterk
```

## Usage
Import the library:
```go
import "github.com/h3th-IV/go-paysterk/paysterk"
```

## Initialize a Transaction
```go
client := paysterk.NewClient("your-secret-key")
tx, err := client.InitializeTransaction("customer@example.com", 5000, "NGN")
if err != nil {
    log.Fatal(err)
}
fmt.Println("Payment URL:", tx.AuthorizationURL)
```

## Verify a Transaction
```go
status, err := client.VerifyTransaction("transaction-reference")
if err != nil {
    log.Fatal(err)
}
fmt.Println("Transaction Status:", status)
```

## Webhook Handler
```go
//The webhook handler handles some of the webhook events supported, this would get real time updates from paystack
http.HandleFunc("/go-paysterk", paysterk.WebHookHandler)
log.Println("webhook server running and listening on :9090")
log.Fatal(http.ListenAndServe(":8080", nil))
```

## Contributing
- Fork the repo
- Create a new branch (feature/new-feature)
- Commit your changes
- Push to your branch
- Open a pull request

## License
This project is licensed under the MIT License.
