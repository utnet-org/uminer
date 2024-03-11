package cmd

import (
	"fmt"
	"google.golang.org/grpc"
)

const WorkerServerIP = "192.168.10.56"
const MinerServerIP = "192.168.10.56"
const AccountId = "jackronwong"

// NodeURL nodeURL
const NodeURL = "http://43.198.88.81:3031"

// LatestBlockHeight latest BlockHeight
var LatestBlockH int64

func ConnectRPCServer(ip string, port string) *grpc.ClientConn {
	conn, err := grpc.Dial(ip+":"+port, grpc.WithInsecure()) // Replace with the correct server address and port
	if err != nil {
		fmt.Println("Error connecting to RPC server:", err)
		return &grpc.ClientConn{}
	}
	return conn
}
