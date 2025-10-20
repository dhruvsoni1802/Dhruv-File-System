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

	//Each chunk is 20KB in sizee, hence based on the file size, we calculate the number of chunks and then read the file using the chunk index
	fileSize := 60 * 1024 // 60KB
	chunkSize := 20 * 1024 // 20KB
	numChunks := fileSize / chunkSize

	//We read the last chunk of the file as an example for now (Chunk indexes are 0-indexed)
	chunkServerAddresses, err := clerk.ReadFile("test.txt", uint64(numChunks - 1))
	if err != nil {
		log.Fatal("Error reading file:", err)
	}
	fmt.Println("Chunk server addresses:", chunkServerAddresses)
}