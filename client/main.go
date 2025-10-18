package main

import (
	"fmt"
	"log"
)

func main() {
	clerk := NewClerk()

	clerk.AddServer("math", "localhost:1234")
	clerk.AddServer("master", "localhost:1235")

	result, err := clerk.Add(3, 2)
	if err != nil {
		log.Fatal("Error adding numbers:", err)
	}
	fmt.Println("Result:", result)

	chunkServerAddresses, err := clerk.ReadFile("test.txt", 0)
	if err != nil {
		log.Fatal("Error reading file:", err)
	}
	fmt.Println("Chunk server addresses:", chunkServerAddresses)

	chunkServerAddresses, err = clerk.WriteFile("test.txt",64)
	if err != nil {
		log.Fatal("Error writing file:", err)
	}
	fmt.Println("Chunk server addresses:", chunkServerAddresses)
}