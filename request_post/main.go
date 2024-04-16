package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type PostRequest struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type PostResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func main() {
	// Create a sample POST request payload
	requestBody, err := json.Marshal(PostRequest{
		Name:  "Apple",
		Price:   1,
	})
	if err != nil {fmt.Printf("Failed to marshal JSON: %v", err)}

	// Send POST request to JSONPlaceholder's /posts endpoint
	resp, err := http.Post("http://127.0.0.1:3000/api/items", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {fmt.Printf("Failed to send POST request: %v", err)}
	defer resp.Body.Close()

	// Decode response body into PostResponse struct
	var postResponse PostResponse
	if err := json.NewDecoder(resp.Body).Decode(&postResponse); err != nil {fmt.Printf("Failed to decode response JSON: %v", err)}

	// Print the created post details
	fmt.Printf("Created Post: ID:%d, Name:%s, Price:%d\n", postResponse.ID, postResponse.Name, postResponse.Price)
}
