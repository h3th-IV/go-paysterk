## go-paysterk
A Golang library for integrating with Paystack's payment API.


## Overview
go-paysterk is a Go client library for interacting with Paystack's API. It simplifies payment processing, transaction verification, transfers, and more.

## Features
- Initialize and verify transactions
- Manage subscriptions and recurring billing
- Handle bank transfers, USSD, and mobile money payments
- Split payments and manage subaccounts
- Process refunds and chargebacks
- Webhook support for real-time transaction updates
- Charge Usage if you dont want redirect users to paystack checkout

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
