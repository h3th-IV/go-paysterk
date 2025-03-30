package paysterk

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// WebHoookhandler process incmoing paystack webhook(notification) request
func WebHookHandler(w http.ResponseWriter, r *http.Request) {
	//Read request Body
	byteBody, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	//Verify the request siganature (security measure)
	secretKey := os.Getenv("PAYSTACK_SECRET_KEY") //this will be replaced
	signature := r.Header.Get("x-paystack-sginature")

	if !isValidSignare(byteBody, signature, secretKey) {
		http.Error(w, "Invalid Signature", http.StatusBadRequest)
		return
	}

	//Parse the JSON reuest body into WebHookEven struct
	var event WebhookEvent
	if err := json.Unmarshal(byteBody, event); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}

	//Process the webhook based on event type
	switch event.Event {
	//Transaction(charge) events
	case "charge.success":
		var transactionData TransactionEventData
		json.Unmarshal(byteBody, &transactionData) // Parse into specific struct
		fmt.Printf("Payment successful: %s (₦%d)\n", transactionData.Data.Customer.Email, transactionData.Data.Amount)

	//Transfer events
	case "transfer.success":
		var transferData TransferEventData
		json.Unmarshal(byteBody, &transferData) // Parse into specific struct
		fmt.Printf("Transfer completed: %s (₦%d)\n", transferData.Data.Recipient.Details.AccountName, transferData.Data.Amount)
	case "transfer.failed":
		fmt.Println("Transfer failed")
	case "transfer.reversed":
		fmt.Println("Transfer reversed")

	//Subscription events
	case "subscription.create":
		var subscriptionData SubscriptionEventData
		json.Unmarshal(byteBody, &subscriptionData) // Parse into specific struct
		fmt.Printf("Subscription created: %s (₦%d)\n", subscriptionData.Data.Customer.Email, subscriptionData.Data.Amount)
	case "subscription.disable":
		fmt.Println("Subscription disabled")
	case "subscription.not_renew":
		fmt.Println("Subscription not renewed")
	case "subscription.expiring_cards":
		fmt.Println("Subscription expiring cards")

	//payment request event
	case "paymentrequest.pending":
		var paymentRequestData PaymentRequestEventData
		json.Unmarshal(byteBody, &paymentRequestData) // Parse into specific struct
		fmt.Printf("Payment request pending: (₦%d)\n", paymentRequestData.Data.Amount)
	case "paymentrequest.success":
		var paymentRequestData PaymentRequestEventData
		json.Unmarshal(byteBody, &paymentRequestData) // Parse into specific struct
		fmt.Printf("Payment request successful: (₦%d)\n", paymentRequestData.Data.Amount)

	//refund events
	case "refund.processed":
		var refundData RefundEventData
		json.Unmarshal(byteBody, &refundData) // Parse into specific struct
		fmt.Printf("Refund successful: %s (₦%d)\n", refundData.Data.Customer.Email, refundData.Data.Amount)
	case "refund.failed":
		fmt.Println("Refund failed")
	case "refund.processing":
		fmt.Println("Refund processing")
	case "refund.pending":
		fmt.Println("Refund initiated, waiting for response")

	//dispute events
	case "charge.dispute.create":
		var disputeData DisputeEventData
		json.Unmarshal(byteBody, &disputeData) // Parse into specific struct
		fmt.Printf("Dispute created: %s (₦%d)\n", disputeData.Data.Customer.Email, disputeData.Data.RefundAmount)
	case "charge.dispute.remind":
		fmt.Println("Dispute reminder: dispute not resolved")
	case "charge.dispute.resolve":
		fmt.Println("Dispute resolved")

	//customer identification
	case "customeridentification.success":
		var identificationData CustomerIdentificationEventData
		json.Unmarshal(byteBody, &identificationData) // Parse into specific struct
		fmt.Printf("Customer identification successful: %s\n", identificationData.Data.Email)
	case "customeridentification.failed":
		fmt.Println("Customer identification failed")
	default:
		fmt.Printf("Unhandled event type: %s\n", event.Event)
	}
}

// isValidSignare checks if th webhook request is from Paystack
func isValidSignare(requestBody []byte, signature, secretKey string) bool {
	if signature == "" {
		return false
	}

	//Create a HMAC-SHA512 hash using the secretKey
	hash := hmac.New(sha512.New, []byte(secretKey))
	hash.Write(requestBody)
	expectedSignature := fmt.Sprintf("%x", hash.Sum(nil))

	//compare request signature and expetected signature
	return hmac.Equal([]byte(signature), []byte(expectedSignature))
}
