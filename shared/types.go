package shared

//This struct is the client input for the Add method
type Args struct {
	A, B int
}

//This struct is the client input for the ReadFile method
type ReadFileArgsMaster struct {
	FileName string
	Offset uint64
}

//This struct is the client input for the WriteFile method
type WriteFileArgsMaster struct {
	FileName string
	DataSize uint64
}

//This struct is the reply from the server for RPC Add method
type MathReply struct {
	Result int
}

//This struct is the reply from the server for RPC ReadFile method
type ReadFileReply struct {
	ChunkServerAddresses []string
}

//This struct is the reply from the server for RPC WriteFile method
type WriteFileReply struct {
	ChunkServerAddresses []string
}