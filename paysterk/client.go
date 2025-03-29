package paysterk

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

// doRequest is used to make HTTP request to various endpoints
func (c *PaystackCLient) doRequest(method, endpoint string, payload interface{}) ([]byte, error) {
	url := fmt.Sprintf("%s/%s", c.BaseURL, endpoint)

	//marshal paylaod to JSON if provided
	var reqBody io.Reader
	if payload != nil {
		jsonData, err := json.Marshal(payload)
		if err != nil {
			return nil, err
		}
		reqBody = bytes.NewBuffer(jsonData)
	}

	//create HTTP request
	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return nil, err
	}

	//Set request Haeaders
	req.Header.Set("Authorization", "Bearer "+c.SecretKey)
	req.Header.Set("Content-Type", "application/json")

	//Make request
	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close() //Ensure response is closed after execution

	//Read response body
	byteBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	//check for Non-2XX status code, so we catch unexpected errors
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		err := fmt.Sprintf("APi error: %v", string(byteBody))
		return nil, errors.New(err)
	}
	return byteBody, nil
}
