package main

import (
	"fmt"
	"log"
	"net/rpc"
)

//This struct is the client input for the Add method
type Args struct {
	A, B int
}

func main() {
	client, err := rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("Error connecting to server:", err)
	}

	//Close the connection when the program exits
	defer client.Close()

	var args Args = Args{A: 1, B: 2}
	var reply int
	err = client.Call("MathService.Add", &args, &reply)
	if err != nil {
		log.Fatal("Error calling Add:", err)
	}
	fmt.Println("Result:", reply)
}