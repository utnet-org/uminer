package cmd

import (
	"fmt"
	"google.golang.org/grpc"
)

// CommanderServerIP IP address of miner (for a single miner deployment)
var CommanderServerIP = "127.0.0.1"

// WorkerServerIP IP address of worker (for a single worker deployment)
var WorkerServerIP = "127.0.0.1"

// NodeURL utility nodeURL
var NodeURL = "http://127.0.0.1:0"

// WorkerLists my worker lists(when deployed on miner server, all available worker is listed below)
var WorkerLists = []string{"127.0.0.1"}

// LatestBlockH marks the latest height of new block
var LatestBlockH int64

func ConnectRPCServer(ip string, port string) *grpc.ClientConn {
	conn, err := grpc.Dial(ip+":"+port, grpc.WithInsecure())
	if err != nil {
		fmt.Println("Error connecting to RPC server:", err)
		return &grpc.ClientConn{}
	}
	return conn
}
