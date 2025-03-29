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

## Installation
```sh
go get github.com/h3th-IV/go-paysterk
```

## Usage
Import the library:
```go
import "github.com/h3th-IV/go-paysterk/paystack"
```

## Initialize a Transaction
```go
client := paystack.NewClient("your-secret-key")
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

## Contributing
- Fork the repo
- Create a new branch (feature/new-feature)
- Commit your changes
- Push to your branch
- Open a pull request

## License
This project is licensed under the MIT License.
