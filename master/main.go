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
type Master struct {}

//This method will be called by MasterClerk to perform a read operation on a particular file
func (m *Master) ReadFile(args *shared.ReadFileArgsMaster, reply *shared.ReadFileReply) error {
	reply.ChunkServerAddresses = []string{"localhost:1234", "localhost:1235", "localhost:1236"}
		return nil
}

//This method will be called by MasterClerk to perform a write operation on a particular file
func (m *Master) WriteFile(args *shared.WriteFileArgsMaster, reply *shared.WriteFileReply) error {
	reply.ChunkServerAddresses = []string{"localhost:1234", "localhost:1235", "localhost:1236"}
	return nil
}

func main() {
	master := new(Master)
	rpc.Register(master)
	rpc.HandleHTTP()
	listener, err := net.Listen("tcp", ":1235") // TCP listening on port 5000
	if err != nil {
		log.Fatal("Error starting RPC server:", err)
	}
	fmt.Println("Server is running on port 1235")
	http.Serve(listener, nil)
}