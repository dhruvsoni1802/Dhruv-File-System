package main

import "sync"

//The master server maintains two data structures:
// 1. A map of file name to chunk handles
// 2. A map of chunk handles to chunk server addresses
type ChunkIndexStore struct{
	Chunk_Index uint64
	Chunk_Server_Addresses []string
}

type FileIndexStore struct {
	File_Name string
	Chunk_Indexes []uint64
}

// Global maps with their respective locks for thread-safe concurrent access
var (
	FileIndexStores map[string]FileIndexStore = make(map[string]FileIndexStore)
	fileIndexMutex  sync.RWMutex  // Protects FileIndexStores map
	
	ChunkIndexStores map[uint64]ChunkIndexStore = make(map[uint64]ChunkIndexStore)
	chunkIndexMutex  sync.RWMutex  // Protects ChunkIndexStores map
)