package master

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
func (m *Master) ReadFile(fileName string, reply *shared.ReadFileReply) error {
	reply.Result = "File read successfully"
		return nil
}

//This method will be called by MasterClerk to perform a write operation on a particular file
func (m *Master) WriteFile(fileName string, data string, reply *shared.WriteFileReply) error {
	reply.Result = "File written successfully"
	return nil
}

func main() {
	master := new(Master)
	rpc.Register(master)
	rpc.HandleHTTP()
	listener, err := net.Listen("tcp", ":5000") // TCP listening on port 5000
	if err != nil {
		log.Fatal("Error starting RPC server:", err)
	}
	fmt.Println("Server is running on port 5000")
	http.Serve(listener, nil)
}