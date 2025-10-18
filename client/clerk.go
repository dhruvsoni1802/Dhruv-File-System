package main

import (
	"DFS/shared"
	"errors"
	"fmt"
	"net/rpc"
)

//This struct is the Clerk that will be used to communicate with the servers
type MathClerk struct {
	clients map[string]*rpc.Client
}

//This method creates a new MathClerk connected to one server
func NewMathClerk(serverAddress string) (*MathClerk, error) {
	client, err := rpc.DialHTTP("tcp", serverAddress)
	if err != nil {
		return nil, err
	}
	
	// Start with one server
	clients := make(map[string]*rpc.Client)
	clients[serverAddress] = client
	
	return &MathClerk{clients: clients}, nil
}

// AddServer adds another replica server to the pool
// Use this to add backup servers
func (c *MathClerk) AddServer(serverAddress string) error {
	client, err := rpc.DialHTTP("tcp", serverAddress)
	if err != nil {
		return err
	}
	c.clients[serverAddress] = client
	return nil
}

// Add sends Add RPC to servers, tries each one until one succeeds
// This is useful when you have replicated servers
func (c *MathClerk) Add(a, b int) (int, error) {
	args := &shared.Args{A: a, B: b}
	var reply shared.Reply

	// Try each server in the map
	for serverAddr, client := range c.clients {
		err := client.Call("MathService.Add", args, &reply)
		
		// If this server succeeded, return the result
		if err == nil {
			fmt.Printf("Successfully called Add on %s\n", serverAddr)
			return reply.Result, nil
		}
		
		// If this server failed, log and try the next one
		fmt.Printf("Server %s failed: %v (trying next...)\n", serverAddr, err)
	}

	// All servers failed
	return 0, errors.New("all servers failed")
}

//To close all connections
func (c *MathClerk) Close() error {
	for serverAddr, client := range c.clients {
		err := client.Close()
		if err != nil {
			fmt.Printf("Error closing connection to %s: %v\n", serverAddr, err)
			return err
		}
	}
	return nil
}