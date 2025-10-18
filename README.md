# Dhruv File System (DFS)

A distributed file system implementation in Go using RPC (Remote Procedure Call) architecture.

## Project Overview

This project implements a distributed file system with the following components:
- **Math Server**: A simple RPC server that performs mathematical operations (demonstration of RPC functionality)
- **Master Server**: Coordinates file operations and manages chunk server addresses
- **Client**: Interacts with both Math and Master servers to perform operations

## Prerequisites

- Go 1.25.1 or higher
- Terminal/Command line access

## Project Structure

```
Dhruv File System/
├── client/          # Client application
│   ├── clerk.go     # Client-side RPC wrapper (Clerk)
│   └── main.go      # Client entry point
├── master/          # Master server
│   └── main.go      # Master server implementation
├── math/            # Math server
│   └── main.go      # Math server implementation
├── shared/          # Shared data structures
│   └── types.go     # Common types and structs
├── go.mod           # Go module file
└── README.md        # This file
```

## Running the System

The system requires all three components to be running simultaneously. Follow these steps in order:

### Step 1: Start the Math Server

Open a terminal window and run:

```bash
cd "/Users/dhruvsoni/Desktop/Distributed systems/Dhruv File System/math"
go run main.go
```

Expected output:
```
Server is running on port 1234
```

The Math server will listen on `localhost:1234`.

### Step 2: Start the Master Server

Open a **new** terminal window and run:

```bash
cd "/Users/dhruvsoni/Desktop/Distributed systems/Dhruv File System/master"
go run main.go
```

Expected output:
```
Server is running on port 1235
```

The Master server will listen on `localhost:1235`.

### Step 3: Run the Client

Open a **third** terminal window and run:

```bash
cd "/Users/dhruvsoni/Desktop/Distributed systems/Dhruv File System/client"
go run client/*.go
```

Expected output:
```
Result: 5
Chunk server addresses: [localhost:1234 localhost:1235 localhost:1236]
Chunk server addresses: [localhost:1234 localhost:1235 localhost:1236]
```

## What the Client Does

The client performs the following operations:

1. **Mathematical Operation**: Adds two numbers (3 + 2) using the Math server via RPC
2. **Read File**: Requests chunk server addresses for reading `test.txt` from the Master server
3. **Write File**: Requests chunk server addresses for writing to `test.txt` (64 bytes) from the Master server

## Architecture

### Math Server
- **Port**: 1234
- **Purpose**: Demonstrates basic RPC functionality with mathematical operations
- **RPC Methods**:
  - `MathService.Add`: Adds two numbers

### Master Server
- **Port**: 1235
- **Purpose**: Coordinates file operations and returns chunk server addresses
- **RPC Methods**:
  - `Master.ReadFile`: Returns chunk server addresses for reading a file
  - `Master.WriteFile`: Returns chunk server addresses for writing to a file

### Client (Clerk)
- **Purpose**: Provides a simplified interface to interact with both servers
- **Features**:
  - Manages RPC connections to multiple servers
  - Provides type-safe methods for server operations

## Development

To modify the project:

1. **Add new RPC methods**: Define them in the respective server files (math/main.go or master/main.go)
2. **Add new data types**: Add shared types to `shared/types.go`
3. **Extend the client**: Modify `client/clerk.go` to add new client-side methods

## Future Enhancements

- Implement actual chunk servers for file storage
- Add persistent storage for file metadata
- Implement file replication and fault tolerance
- Add authentication and authorization
- Implement actual file read/write operations

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

This is an educational project for learning distributed systems concepts.

