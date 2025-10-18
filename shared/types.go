package shared

//This struct is the client input for the Add method
type Args struct {
	A, B int
}

//This struct is the reply from the server for RPC Add method
type Reply struct {
	Result int
}