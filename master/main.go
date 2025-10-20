package main

import (
	"DFS/shared"
	"errors"
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
	//Acquire a READ lock - multiple readers can hold this simultaneously
	chunkIndexMutex.RLock()
	//Now we look at the chunk index from the args and fetch the chunk server addresses from the ChunkIndexStores
	chunkIndexStore, exists := ChunkIndexStores[args.Chunk_Index]
	//Release the READ lock immediately after reading
	chunkIndexMutex.RUnlock()
	
	if !exists {
		return errors.New("chunk index not found")
	}
	reply.ChunkServerAddresses = chunkIndexStore.Chunk_Server_Addresses
	return nil
}

//This method will be called by MasterClerk to perform a write operation on a particular file
func (m *Master) WriteFile(args *shared.WriteFileArgsMaster, reply *shared.WriteFileReply) error {
	reply.ChunkServerAddresses = []string{"localhost:1234", "localhost:1235", "localhost:1236"}
	return nil
}

func main() {
	//For now because the write path is not implemented,
	//We will populate the data structures with some dummy data to test the read path
	
	//Acquire WRITE lock for FileIndexStores - exclusive access for writing
	fileIndexMutex.Lock()
	FileIndexStores["test.txt"] = FileIndexStore{
		File_Name: "test.txt",
		Chunk_Indexes: []uint64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
	}
	fileIndexMutex.Unlock()

	//Acquire WRITE lock for ChunkIndexStores - exclusive access for writing
	chunkIndexMutex.Lock()
	//Assuming a replication factor of 3 for now
	ChunkIndexStores[0] = ChunkIndexStore{
		Chunk_Index: 0,
		Chunk_Server_Addresses: []string{"localhost:5001", "localhost:5002", "localhost:5003"},
	}

	ChunkIndexStores[1] = ChunkIndexStore{
		Chunk_Index: 1,
		Chunk_Server_Addresses: []string{"localhost:5002", "localhost:5003", "localhost:5004"},
	}

	ChunkIndexStores[2] = ChunkIndexStore{
		Chunk_Index: 2,
		Chunk_Server_Addresses: []string{"localhost:5003", "localhost:5004", "localhost:5001"},
	}
	chunkIndexMutex.Unlock()

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