package main

import (
	"DFS/shared"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

//This struct is the receiver for the RPC methods
type MathService struct{}

//This method will be called by MathClerk to add two numbers via RPC
func (m *MathService) Add(args *shared.Args, reply *shared.MathReply) error {
	reply.Result = args.A + args.B
	return nil
}

func main() {
	mathService := new(MathService)
	rpc.Register(mathService)
	rpc.HandleHTTP()
	listener, err := net.Listen("tcp", ":1234") // TCP listening on port 1234
	if err != nil {
		log.Fatal("Error starting RPC server:", err)
	}
	fmt.Println("Server is running on port 1234")
	http.Serve(listener, nil)
}

