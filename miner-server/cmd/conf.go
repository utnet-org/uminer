package cmd

import (
	"fmt"
	"google.golang.org/grpc"
)

// MinerServerIP IP address of miner (for a single miner deployment)
const MinerServerIP = "192.168.10.56"

// WorkerServerIP IP address of worker (for a single worker deployment)
const WorkerServerIP = "192.168.10.56"

// NodeURL utility nodeURL
const NodeURL = "http://43.198.88.81:3031"

// WorkerLists my worker lists(when deployed on miner server, all available worker is listed below)
var WorkerLists = [...]string{"192.168.10.50", "192.168.10.56", "192.168.10.59"}

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
