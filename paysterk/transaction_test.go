package paysterk

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// NOTE: all keys and references passed here are for test purpose only
// Mock HTTP server
func setupMockServer(response string, statusCode int) *httptest.Server {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(statusCode)
		w.Write([]byte(response))
	})
	return httptest.NewServer(handler)
}

func TestInitializeTransaction(t *testing.T) {
	mockServer := setupMockServer(`{"status": true, "message": "Success", "data": {"authorization_url": "https://paystack.com/pay/xyz"}}`, 200)
	defer mockServer.Close()

	client := &PaystackCLient{SecretKey: "test_key", BaseURL: mockServer.URL}

	tx, err := client.InitializeTransaction("user@example.com", 5000, "NGN")
	if err != nil || tx.Data.AuthorizationURL == "" {
		t.Errorf("Failed to initialize transaction")
	}
}

func TestVerifyTransaction(t *testing.T) {
	mockServer := setupMockServer(`{"status": true, "message": "Verification successful", "data": {"status": "success"}}`, 200)
	defer mockServer.Close()

	client := &PaystackCLient{SecretKey: "test_key", BaseURL: mockServer.URL}

	resp, err := client.VerifyTransaction("test_ref")
	if err != nil || resp.Data.Status != "success" {
		t.Errorf("Failed to verify transaction")
	}
}

// Test fetching all transactions with pagination
func TestFetchAllTransactions(t *testing.T) {
	mockResponse := `{
		"status": true,
		"message": "Transactions retrieved",
		"data": [
			{
				"id": 4099260516,
				"reference": "re4lyvq3s3",
				"amount": 40333,
				"currency": "NGN",
				"status": "success",
				"channel": "card",
				"paidAt": "2024-08-22T09:15:02.000Z",
				"createdAt": "2024-08-22T09:14:24.000Z",
				"gateway_response": "Successful",
				"fees": 10283,
				"customer": {
					"id": 181873746,
					"email": "demo@test.com",
					"customer_code": "CUS_1rkzaqsv4rrhqo6"
				}
			}
		],
		"meta": {
			"next": "dW5kZWZpbmVkOjQwMTM3MDk2MzU=",
			"previous": null,
			"perPage": 50
		}
	}`

	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(mockResponse))
	}))
	defer mockServer.Close()

	client := &PaystackCLient{SecretKey: "test_key", BaseURL: mockServer.URL}

	response, err := client.FetchAllTransactions()
	if err != nil {
		t.Errorf("Error fetching transactions: %v", err)
	}

	if response.Status != true {
		t.Errorf("Expected status to be true, got %v", response.Status)
	}

	if len(response.Data) != 1 {
		t.Errorf("Expected 1 transaction, got %d", len(response.Data))
	}

	if response.Meta.PerPage != 50 {
		t.Errorf("Expected perPage to be 50, got %d", response.Meta.PerPage)
	}
}

// Test fetching a single transaction by ID
func TestFetchTransactionByID(t *testing.T) {
	mockResponse := `{
		"status": true,
		"message": "Transaction retrieved",
		"data": {
			"id": 4099260516,
			"reference": "re4lyvq3s3",
			"amount": 40333,
			"currency": "NGN",
			"status": "success",
			"channel": "card",
			"paidAt": "2024-08-22T09:15:02.000Z",
			"createdAt": "2024-08-22T09:14:24.000Z",
			"gateway_response": "Successful",
			"fees": 10283,
			"customer": {
				"id": 181873746,
				"email": "demo@test.com",
				"customer_code": "CUS_1rkzaqsv4rrhqo6"
			}
		}
	}`

	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(mockResponse))
	}))
	defer mockServer.Close()

	client := &PaystackCLient{SecretKey: "test_key", BaseURL: mockServer.URL}

	transaction, err := client.FetchTransactionByID(4099260516)
	if err != nil {
		t.Errorf("Error fetching transaction: %v", err)
	}

	if transaction.ID != 4099260516 {
		t.Errorf("Expected transaction ID 4099260516, got %d", transaction.ID)
	}

	if transaction.Status != "success" {
		t.Errorf("Expected status 'success', got %s", transaction.Status)
	}
}
