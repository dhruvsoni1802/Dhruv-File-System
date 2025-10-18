package main

import (
	"fmt"
	"log"
)

func main() {
	client, err := NewMathClerk("localhost:1234")
	if err != nil {
		log.Fatal("Error creating MathClerk:", err)
	}
	defer client.Close()

	result, err := client.Add(3, 2)
	if err != nil {
		log.Fatal("Error adding numbers:", err)
	}
	fmt.Println("Result:", result)
}