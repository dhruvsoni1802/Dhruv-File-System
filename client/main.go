package main

import (
	"fmt"
	"log"
)

func main() {
	clerk := NewClerk()

	clerk.AddServer("math", "localhost:1234")

	result, err := clerk.Add(3, 2)
	if err != nil {
		log.Fatal("Error adding numbers:", err)
	}
	fmt.Println("Result:", result)
}