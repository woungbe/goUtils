package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Identifications struct {
	Identifications []Chainalysis `json:"identifications"`
}

type Chainalysis struct {
	Category    string `json:"category" db:"category"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	Url         string `json:"url" db:"url"`
}

func main() {

	address := "0x1da5821544e25c636c1417ba96ade4cf6d2f9b5a"
	// address := "0x2ea58f9858d7f5f150a562c09ab6532339918b90"//
	// url := "https://public.chainalysis.com/api/v1/address/0x2ea58f9858d7f5f150a562c09ab6532339918b90"
	url := fmt.Sprintf("https://public.chainalysis.com/api/v1/address/%s", address)

	// Create a new request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set headers
	req.Header.Set("X-API-Key", "be9ea4af8edce390dda56500d4e046665f435e7b90dc659fdf630ba9d17f0d12")
	req.Header.Set("Accept", "application/json")

	// Create a new HTTP client and perform the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error performing request:", err)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// Print the response body
	fmt.Println(string(body))

	var iden Identifications
	err = json.Unmarshal(body, &iden)
	if err != nil {
		fmt.Println(err)
	}

	for _, v := range iden.Identifications {
		fmt.Printf("Url: %s\n Name:%s\n, Category:%s\n, Description:%s\n", v.Url, v.Name, v.Category, v.Description)
	}

}
