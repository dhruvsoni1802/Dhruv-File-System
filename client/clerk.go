package main

import (
	"DFS/shared"
	"errors"
	"fmt"
	"net/rpc"
)

// Clerk struct with a generic map of server types, each containing their own client pools
type Clerk struct {
	servers map[string]map[string]*rpc.Client // servers[serverType][address] = client
}

// NewClerk creates a new Clerk
func NewClerk() *Clerk {
	return &Clerk{
		servers: make(map[string]map[string]*rpc.Client),
	}
}

// AddServer adds a new server of a given type to the clerk's client pool
func (c *Clerk) AddServer(serverType, serverAddress string) error {
	client, err := rpc.DialHTTP("tcp", serverAddress)
	if err != nil {
		return err
	}
	
	// Initialize the map for this server type if it doesn't exist
	if _, exists := c.servers[serverType]; !exists {
		c.servers[serverType] = make(map[string]*rpc.Client)
	}
	
	c.servers[serverType][serverAddress] = client
	return nil
}

// RemoveServer removes a server of a given type from the clerk's client pool
func (c *Clerk) RemoveServer(serverType, serverAddress string) error {
	clients, exists := c.servers[serverType]
	if !exists {
		return errors.New("server type not found")
	}
	
	client, exists := clients[serverAddress]
	if !exists {
		return errors.New("server address not found")
	}
	
	err := client.Close()
	if err != nil {
		return err
	}
	
	delete(clients, serverAddress)
	return nil
}

// callRPC is a generic helper that tries an RPC call on each server of a given type until one succeeds
func (c *Clerk) callRPC(serverType, method string, args interface{}, reply interface{}) error {
	clients, exists := c.servers[serverType]
	if !exists {
		return errors.New("no servers of type: " + serverType)
	}
	
	if len(clients) == 0 {
		return errors.New("no available servers of type: " + serverType)
	}
	
	for serverAddr, client := range clients {
		err := client.Call(method, args, reply)
		
		if err == nil {
			fmt.Printf("Successfully called %s on %s (%s server)\n", method, serverAddr, serverType)
			return nil
		}
		
		fmt.Printf("Server %s failed: %v (trying next...)\n", serverAddr, err)
	}
	
	return errors.New("all " + serverType + " servers failed")
}

// Add sends Add RPC to math servers, tries each one until one succeeds
func (c *Clerk) Add(a, b int) (int, error) {
	args := &shared.Args{A: a, B: b}
	var reply shared.MathReply
	
	err := c.callRPC("math", "MathService.Add", args, &reply)
	if err != nil {
		return 0, err
	}
	
	return reply.Result, nil
}

// ReadFile sends ReadFile RPC to master servers
func (c *Clerk) ReadFile(fileName string, offset uint64) ([]string, error) {
	args := &shared.ReadFileArgsMaster{FileName: fileName, Offset: offset}
	var reply shared.ReadFileReply
	
	err := c.callRPC("master", "Master.ReadFile", args, &reply)
	if err != nil {
		return nil, err
	}
	
	return reply.ChunkServerAddresses, nil
}

// WriteFile sends WriteFile RPC to master servers
func (c *Clerk) WriteFile(fileName string, dataSize uint64) ([]string, error) {
	args := &shared.WriteFileArgsMaster{FileName: fileName, DataSize: dataSize}
	var reply shared.WriteFileReply
	
	err := c.callRPC("master", "Master.WriteFile", args, &reply)
	if err != nil {
		return nil, err
	}
	
	return reply.ChunkServerAddresses, nil
}


func (c *Clerk) Close() error {
	for serverType, clients := range c.servers {
		for serverAddr, client := range clients {
			err := client.Close()
			if err != nil {
				fmt.Printf("Error closing connection to %s server %s: %v\n", serverType, serverAddr, err)
				return err
			}
		}
	}
	return nil
}