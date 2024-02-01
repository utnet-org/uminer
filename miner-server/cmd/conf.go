package cmd

import (
	"fmt"
	"google.golang.org/grpc"
)

const WorkerServerIP = "192.168.10.49"
const MinerServerIP = "192.168.10.49"

func ConnectRPCServer(ip string, port string) *grpc.ClientConn {
	conn, err := grpc.Dial(ip+":"+port, grpc.WithInsecure()) // Replace with the correct server address and port
	if err != nil {
		fmt.Println("Error connecting to RPC server:", err)
		return &grpc.ClientConn{}
	}
	return conn
}
