package main

import (
	"fmt"
	"log"

	"github.com/jeffotoni/quick/http/client"
)

func main() {
	// Define a struct to send as JSON
	data := struct {
		Message string `json:"message"`
	}{
		Message: "Hello, POST Quick New!!!",
	}
	resp, err := client.Put("http://localhost:3000/v1/user/1234", data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("PUT response:", string(resp.Body))
}
