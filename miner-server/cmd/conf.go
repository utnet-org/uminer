package cmd

import (
	"fmt"
	"google.golang.org/grpc"
)

const ServerIP = "192.168.10.49"

func ConnectRPCServer() *grpc.ClientConn {
	conn, err := grpc.Dial(ServerIP+":9001", grpc.WithInsecure()) // Replace with the correct server address and port
	if err != nil {
		fmt.Println("Error connecting to RPC server:", err)
		return &grpc.ClientConn{}
	}
	return conn
}
